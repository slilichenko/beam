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

// Package extworker provides an external worker service and related utilities.
package extworker

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/apache/beam/sdks/v2/go/container/tools"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/harness"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	fnpb "github.com/apache/beam/sdks/v2/go/pkg/beam/model/fnexecution_v1"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/util/grpcx"
	"google.golang.org/grpc"
)

// StartLoopback initializes a Loopback ExternalWorkerService, at the given port.
func StartLoopback(ctx context.Context, port int) (*Loopback, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, err
	}

	log.Infof(ctx, "starting Loopback server at %v", lis.Addr())
	grpcServer := grpc.NewServer()
	root, cancel := context.WithCancel(ctx)
	s := &Loopback{lis: lis, root: root, rootCancel: cancel, workers: map[string]context.CancelFunc{},
		grpcServer: grpcServer}
	fnpb.RegisterBeamFnExternalWorkerPoolServer(grpcServer, s)
	go grpcServer.Serve(lis)
	return s, nil
}

// Loopback implements fnpb.BeamFnExternalWorkerPoolServer
type Loopback struct {
	fnpb.UnimplementedBeamFnExternalWorkerPoolServer

	lis        net.Listener
	root       context.Context
	rootCancel context.CancelFunc

	mu      sync.Mutex
	workers map[string]context.CancelFunc

	grpcServer *grpc.Server
}

// StartWorker initializes a new worker harness, implementing BeamFnExternalWorkerPoolServer.StartWorker.
func (s *Loopback) StartWorker(ctx context.Context, req *fnpb.StartWorkerRequest) (*fnpb.StartWorkerResponse, error) {
	log.Debugf(ctx, "starting worker %v", req.GetWorkerId())
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.workers == nil {
		return fnpb.StartWorkerResponse_builder{
			Error: "worker pool shutting down",
		}.Build(), nil
	}

	if _, ok := s.workers[req.GetWorkerId()]; ok {
		return fnpb.StartWorkerResponse_builder{
			Error: fmt.Sprintf("worker with ID %q already exists", req.GetWorkerId()),
		}.Build(), nil
	}
	if req.GetLoggingEndpoint() == nil {
		return fnpb.StartWorkerResponse_builder{Error: fmt.Sprintf("Missing logging endpoint for worker %v", req.GetWorkerId())}.Build(), nil
	}
	if req.GetControlEndpoint() == nil {
		return fnpb.StartWorkerResponse_builder{Error: fmt.Sprintf("Missing control endpoint for worker %v", req.GetWorkerId())}.Build(), nil
	}
	if req.GetLoggingEndpoint().HasAuthentication() || req.GetControlEndpoint().HasAuthentication() {
		return fnpb.StartWorkerResponse_builder{Error: "[BEAM-10610] Secure endpoints not supported."}.Build(), nil
	}

	ctx = grpcx.WriteWorkerID(s.root, req.GetWorkerId())
	ctx, s.workers[req.GetWorkerId()] = context.WithCancel(ctx)

	opts := harnessOptions(ctx, req.GetProvisionEndpoint().GetUrl())

	go harness.MainWithOptions(ctx, req.GetLoggingEndpoint().GetUrl(), req.GetControlEndpoint().GetUrl(), opts)
	return &fnpb.StartWorkerResponse{}, nil
}

func harnessOptions(ctx context.Context, endpoint string) harness.Options {
	var opts harness.Options
	if endpoint == "" {
		return opts
	}
	info, err := tools.ProvisionInfo(ctx, endpoint)
	if err != nil {
		log.Debugf(ctx, "error talking to provision service worker, using defaults: %v", err)
		return opts
	}

	opts.StatusEndpoint = info.GetStatusEndpoint().GetUrl()
	opts.RunnerCapabilities = info.GetRunnerCapabilities()
	return opts
}

// StopWorker terminates a worker harness, implementing BeamFnExternalWorkerPoolServer.StopWorker.
func (s *Loopback) StopWorker(ctx context.Context, req *fnpb.StopWorkerRequest) (*fnpb.StopWorkerResponse, error) {
	log.Infof(ctx, "stopping worker %v", req.GetWorkerId())
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.workers == nil {
		// Worker pool is already shutting down, so no action is needed.
		return &fnpb.StopWorkerResponse{}, nil
	}
	if cancelfn, ok := s.workers[req.GetWorkerId()]; ok {
		cancelfn()
		delete(s.workers, req.GetWorkerId())
		return &fnpb.StopWorkerResponse{}, nil
	}
	return fnpb.StopWorkerResponse_builder{
		Error: fmt.Sprintf("no worker with id %q running", req.GetWorkerId()),
	}.Build(), nil

}

// Stop terminates the service and stops all workers.
func (s *Loopback) Stop(ctx context.Context) error {
	s.mu.Lock()

	log.Debugf(ctx, "stopping Loopback, and %d workers", len(s.workers))
	s.workers = nil
	s.rootCancel()

	// There can be a deadlock between the StopWorker RPC and GracefulStop
	// which waits for all RPCs to finish, so it must be outside the critical section.
	s.mu.Unlock()

	s.grpcServer.GracefulStop()
	return nil
}

// EnvironmentConfig returns the environment config for this service instance.
func (s *Loopback) EnvironmentConfig(context.Context) string {
	return fmt.Sprintf("localhost:%d", s.lis.Addr().(*net.TCPAddr).Port)
}
