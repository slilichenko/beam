// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package artifact

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/internal/errors"
	jobpb "github.com/apache/beam/sdks/v2/go/pkg/beam/model/jobmanagement_v1"
	pipepb "github.com/apache/beam/sdks/v2/go/pkg/beam/model/pipeline_v1"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/util/grpcx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// TestRetrieve tests that we can successfully retrieve fresh files.
func TestRetrieve(t *testing.T) {
	cc := startServer(t)
	defer cc.Close()

	ctx := grpcx.WriteWorkerID(context.Background(), "idA")
	keys := []string{"foo", "bar", "baz/baz/baz"}
	st := "whatever"
	rt, artifacts := populate(ctx, cc, t, keys, 300, st)

	dst := makeTempDir(t)
	defer os.RemoveAll(dst)

	client := jobpb.NewLegacyArtifactRetrievalServiceClient(cc)
	for _, a := range artifacts {
		filename := makeFilename(dst, a.GetName())
		if err := Retrieve(ctx, client, a, rt, dst); err != nil {
			t.Errorf("failed to retrieve %v: %v", a.GetName(), err)
			continue
		}
		verifySHA256(t, filename, a.GetSha256())
	}
}

// TestMultiRetrieve tests that we can successfully retrieve fresh files
// concurrently.
func TestMultiRetrieve(t *testing.T) {
	cc := startServer(t)
	defer cc.Close()

	ctx := grpcx.WriteWorkerID(context.Background(), "idB")
	keys := []string{"1", "2", "3", "4", "a/5", "a/6", "a/7", "a/8", "a/a/9", "a/a/10", "a/b/11", "a/b/12"}
	st := "whatever"
	rt, artifacts := populate(ctx, cc, t, keys, 300, st)

	dst := makeTempDir(t)
	defer os.RemoveAll(dst)

	client := jobpb.NewLegacyArtifactRetrievalServiceClient(cc)
	if err := LegacyMultiRetrieve(ctx, client, 10, artifacts, rt, dst); err != nil {
		t.Errorf("failed to retrieve: %v", err)
	}

	for _, a := range artifacts {
		verifySHA256(t, makeFilename(dst, a.GetName()), a.GetSha256())
	}
}

// populate stages a set of artifacts with the given keys, each with
// slightly different sizes and chucksizes.
func populate(ctx context.Context, cc *grpc.ClientConn, t *testing.T, keys []string, size int, st string) (string, []*jobpb.ArtifactMetadata) {
	scl := jobpb.NewLegacyArtifactStagingServiceClient(cc)

	var artifacts []*jobpb.ArtifactMetadata
	for i, key := range keys {
		a := stage(ctx, scl, t, key, size+7*i, 97+i, st)
		artifacts = append(artifacts, a)
	}
	token, err := Commit(ctx, scl, artifacts, st)
	if err != nil {
		t.Fatalf("failed to commit manifest: %v", err)
		return "", nil
	}
	return token, artifacts
}

// stage stages an artifact with the given key, size and chuck size. The content is
// always 'z's.
func stage(ctx context.Context, scl jobpb.LegacyArtifactStagingServiceClient, t *testing.T, key string, size, chunkSize int, st string) *jobpb.ArtifactMetadata {
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = 'z'
	}

	sha256W := sha256.New()
	sha256W.Write(data)
	hash := hex.EncodeToString(sha256W.Sum(nil))
	md := makeArtifact(key, hash)
	pmd := jobpb.PutArtifactMetadata_builder{
		Metadata:            md,
		StagingSessionToken: st,
	}.Build()

	stream, err := scl.PutArtifact(ctx)
	if err != nil {
		t.Fatalf("put failed: %v", err)
	}
	header := jobpb.PutArtifactRequest_builder{
		Metadata: proto.ValueOrDefault(pmd),
	}.Build()
	if err := stream.Send(header); err != nil {
		t.Fatalf("send header failed: %v", err)
	}

	for i := 0; i < size; i += chunkSize {
		end := i + chunkSize
		if size < end {
			end = size
		}

		chunk := jobpb.PutArtifactRequest_builder{
			Data: jobpb.ArtifactChunk_builder{
				Data: data[i:end],
			}.Build(),
		}.Build()
		if err := stream.Send(chunk); err != nil {
			t.Fatalf("send chunk[%v:%v] failed: %v", i, end, err)
		}
	}
	if _, err := stream.CloseAndRecv(); err != nil {
		t.Fatalf("close failed: %v", err)
	}
	return md
}

// Test for new artifact retrieval.

func TestNewRetrieveWithManyFiles(t *testing.T) {
	expected := map[string]string{"a.txt": "a", "b.txt": "bbb", "c.txt": "cccccccc"}

	client := &fakeRetrievalService{
		artifacts: expected,
	}

	dest := makeTempDir(t)
	defer os.RemoveAll(dest)
	ctx := grpcx.WriteWorkerID(context.Background(), "worker")

	mds, err := newMaterializeWithClient(ctx, client, client.resolvedArtifactsWithStagingTo(), dest)
	if err != nil {
		t.Fatalf("materialize failed: %v", err)
	}

	checkStagedFiles(mds, dest, expected, t)
}

func TestNewRetrieveWithFileGeneratedStageName(t *testing.T) {
	expected := map[string]string{"a.txt": "", "b.txt": "", "c.txt": ""}

	client := &fakeRetrievalService{
		artifacts: expected,
	}

	dest := makeTempDir(t)
	defer os.RemoveAll(dest)
	ctx := grpcx.WriteWorkerID(context.Background(), "worker")

	mds, err := newMaterializeWithClient(ctx, client, client.fileArtifactsWithoutStagingTo(), dest)
	if err != nil {
		t.Fatalf("materialize failed: %v", err)
	}

	generated := make(map[string]string)
	for _, md := range mds {
		name, _ := MustExtractFilePayload(md)
		payload, _ := proto.Marshal(pipepb.ArtifactStagingToRolePayload_builder{
			StagedName: name}.Build())
		generated[name] = string(payload)
	}

	checkStagedFiles(mds, dest, generated, t)
}

func TestNewRetrieveWithUrlGeneratedStageName(t *testing.T) {
	expected := map[string]string{"a.txt": "", "b.txt": "", "c.txt": ""}

	client := &fakeRetrievalService{
		artifacts: expected,
	}

	dest := makeTempDir(t)
	defer os.RemoveAll(dest)
	ctx := grpcx.WriteWorkerID(context.Background(), "worker")

	mds, err := newMaterializeWithClient(ctx, client, client.urlArtifactsWithoutStagingTo(), dest)
	if err != nil {
		t.Fatalf("materialize failed: %v", err)
	}

	generated := make(map[string]string)
	for _, md := range mds {
		name, _ := MustExtractFilePayload(md)
		payload, _ := proto.Marshal(pipepb.ArtifactStagingToRolePayload_builder{
			StagedName: name}.Build())
		generated[name] = string(payload)
	}

	checkStagedFiles(mds, dest, generated, t)
}

func TestNewRetrieveWithSubdir(t *testing.T) {
	expected := map[string]string{"subdir/path/a.txt": "a"}

	client := &fakeRetrievalService{
		artifacts: expected,
	}

	dest := makeTempDir(t)
	defer os.RemoveAll(dest)
	ctx := grpcx.WriteWorkerID(context.Background(), "worker")

	mds, err := newMaterializeWithClient(ctx, client, client.resolvedArtifactsWithStagingTo(), dest)
	if err != nil {
		t.Fatalf("materialize failed: %v", err)
	}

	checkStagedFiles(mds, dest, expected, t)
}

func TestNewRetrieveWithResolution(t *testing.T) {
	expected := map[string]string{"a.txt": "a", "b.txt": "bbb", "c.txt": "cccccccc"}

	client := &fakeRetrievalService{
		artifacts: expected,
	}

	dest := makeTempDir(t)
	defer os.RemoveAll(dest)
	ctx := grpcx.WriteWorkerID(context.Background(), "worker")

	mds, err := newMaterializeWithClient(ctx, client, client.unresolvedArtifacts(), dest)
	if err != nil {
		t.Fatalf("materialize failed: %v", err)
	}

	checkStagedFiles(mds, dest, expected, t)
}

func checkStagedFiles(mds []*pipepb.ArtifactInformation, dest string, expected map[string]string, t *testing.T) {
	if len(mds) != len(expected) {
		t.Errorf("wrong number of artifacts staged %v vs %v", len(mds), len(expected))
	}
	for _, md := range mds {
		name, _ := MustExtractFilePayload(md)
		filename := filepath.Join(dest, filepath.FromSlash(name))
		fd, err := os.Open(filename)
		if err != nil {
			t.Errorf("error opening file %v", err)
		}
		defer fd.Close()

		data := make([]byte, 1<<20)
		n, err := fd.Read(data)
		if err != nil {
			t.Errorf("error reading file %v", err)
		}

		if string(data[:n]) != expected[name] {
			t.Errorf("missmatched contents for %v: '%s' vs '%s'", name, string(data[:n]), expected[name])
		}
	}
}

type fakeRetrievalService struct {
	artifacts map[string]string // name -> content
}

func (fake *fakeRetrievalService) resolvedArtifactsWithStagingTo() []*pipepb.ArtifactInformation {
	var artifacts []*pipepb.ArtifactInformation
	for name, contents := range fake.artifacts {
		payload, _ := proto.Marshal(pipepb.ArtifactStagingToRolePayload_builder{
			StagedName: name}.Build())
		artifacts = append(artifacts, pipepb.ArtifactInformation_builder{
			TypeUrn:     "resolved",
			TypePayload: []byte(contents),
			RoleUrn:     URNStagingTo,
			RolePayload: payload,
		}.Build())
	}
	return artifacts
}

func (fake *fakeRetrievalService) fileArtifactsWithoutStagingTo() []*pipepb.ArtifactInformation {
	var artifacts []*pipepb.ArtifactInformation
	for name := range fake.artifacts {
		payload, _ := proto.Marshal(pipepb.ArtifactFilePayload_builder{
			Path: filepath.Join("/tmp", name)}.Build())
		artifacts = append(artifacts, pipepb.ArtifactInformation_builder{
			TypeUrn:     URNFileArtifact,
			TypePayload: payload,
		}.Build())
	}
	return artifacts
}

func (fake *fakeRetrievalService) urlArtifactsWithoutStagingTo() []*pipepb.ArtifactInformation {
	var artifacts []*pipepb.ArtifactInformation
	for name := range fake.artifacts {
		payload, _ := proto.Marshal(pipepb.ArtifactUrlPayload_builder{
			Url: path.Join("gs://tmp", name)}.Build())
		artifacts = append(artifacts, pipepb.ArtifactInformation_builder{
			TypeUrn:     URNUrlArtifact,
			TypePayload: payload,
		}.Build())
	}
	return artifacts
}

func (fake *fakeRetrievalService) unresolvedArtifacts() []*pipepb.ArtifactInformation {
	return []*pipepb.ArtifactInformation{
		pipepb.ArtifactInformation_builder{
			TypeUrn: "unresolved",
		}.Build(),
	}
}

func (fake *fakeRetrievalService) ResolveArtifacts(ctx context.Context, request *jobpb.ResolveArtifactsRequest, opts ...grpc.CallOption) (*jobpb.ResolveArtifactsResponse, error) {
	response := &jobpb.ResolveArtifactsResponse{}
	for _, dep := range request.GetArtifacts() {
		if dep.GetTypeUrn() == "unresolved" {
			response.SetReplacements(append(response.GetReplacements(), fake.resolvedArtifactsWithStagingTo()...))
		} else {
			response.SetReplacements(append(response.GetReplacements(), dep))
		}
	}
	return response, nil
}

func (fake *fakeRetrievalService) GetArtifact(ctx context.Context, request *jobpb.GetArtifactRequest, opts ...grpc.CallOption) (jobpb.ArtifactRetrievalService_GetArtifactClient, error) {
	switch request.GetArtifact().GetTypeUrn() {
	case "resolved":
		return &fakeGetArtifactResponseStream{data: request.GetArtifact().GetTypePayload()}, nil
	case URNFileArtifact, URNUrlArtifact:
		return &fakeGetArtifactResponseStream{data: request.GetArtifact().GetRolePayload()}, nil
	default:
		return nil, errors.Errorf("Unsupported artifact %v", request.GetArtifact())
	}
}

type fakeGetArtifactResponseStream struct {
	data  []byte
	index int
}

func (fake *fakeGetArtifactResponseStream) Recv() (*jobpb.GetArtifactResponse, error) {
	if fake.index < len(fake.data) {
		fake.index++
		return jobpb.GetArtifactResponse_builder{Data: fake.data[fake.index-1 : fake.index]}.Build(), nil
	}
	return nil, io.EOF
}

func (fake *fakeGetArtifactResponseStream) RecvMsg(any) error {
	return nil
}

func (fake *fakeGetArtifactResponseStream) SendMsg(any) error {
	return nil
}

func (fake *fakeGetArtifactResponseStream) Header() (metadata.MD, error) {
	return nil, nil
}

func (fake *fakeGetArtifactResponseStream) Trailer() metadata.MD {
	return nil
}

func (fake *fakeGetArtifactResponseStream) Context() context.Context {
	return context.Background()
}

func (fake *fakeGetArtifactResponseStream) CloseSend() error {
	return nil
}

func verifySHA256(t *testing.T, filename, hash string) {
	actual, err := computeSHA256(filename)
	if err != nil {
		t.Errorf("failed to compute hash for %v: %v", filename, err)
		return
	}
	if actual != hash {
		t.Errorf("file %v has bad SHA256: %v, want %v", filename, actual, hash)
	}
}

func makeTempDir(t *testing.T) string {
	dir, err := os.MkdirTemp("", "artifact_test_")
	if err != nil {
		t.Errorf("Test failure: cannot create temporary directory: %+v", err)
	}
	return dir
}

func makeTempFiles(t *testing.T, dir string, keys []string, size int) []string {
	var sha256s []string
	for i, key := range keys {
		hash := makeTempFile(t, makeFilename(dir, key), size+i)
		sha256s = append(sha256s, hash)
	}
	return sha256s
}

func makeTempFile(t *testing.T, filename string, size int) string {
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = 'z'
	}

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		t.Fatalf("cannot create directory for %s: %v", filename, err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		t.Fatalf("cannot create file %s: %v", filename, err)
	}

	sha256W := sha256.New()
	sha256W.Write(data)
	return hex.EncodeToString(sha256W.Sum(nil))
}

func makeArtifact(key, hash string) *jobpb.ArtifactMetadata {
	return jobpb.ArtifactMetadata_builder{
		Name:        key,
		Sha256:      hash,
		Permissions: 0644,
	}.Build()
}

func makeFilename(dir, key string) string {
	return filepath.Join(dir, filepath.FromSlash(key))
}
