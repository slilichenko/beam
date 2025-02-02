//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// Protocol Buffers describing endpoints containing a service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.27.3
// source: org/apache/beam/model/pipeline/v1/endpoints.proto

//go:build protoopaque

package pipeline_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A description of how to connect to a Beam API endpoint.
type ApiServiceDescriptor struct {
	state                     protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Url            string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	xxx_hidden_Authentication *AuthenticationSpec    `protobuf:"bytes,2,opt,name=authentication,proto3" json:"authentication,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *ApiServiceDescriptor) Reset() {
	*x = ApiServiceDescriptor{}
	mi := &file_org_apache_beam_model_pipeline_v1_endpoints_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiServiceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiServiceDescriptor) ProtoMessage() {}

func (x *ApiServiceDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_pipeline_v1_endpoints_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ApiServiceDescriptor) GetUrl() string {
	if x != nil {
		return x.xxx_hidden_Url
	}
	return ""
}

func (x *ApiServiceDescriptor) GetAuthentication() *AuthenticationSpec {
	if x != nil {
		return x.xxx_hidden_Authentication
	}
	return nil
}

func (x *ApiServiceDescriptor) SetUrl(v string) {
	x.xxx_hidden_Url = v
}

func (x *ApiServiceDescriptor) SetAuthentication(v *AuthenticationSpec) {
	x.xxx_hidden_Authentication = v
}

func (x *ApiServiceDescriptor) HasAuthentication() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Authentication != nil
}

func (x *ApiServiceDescriptor) ClearAuthentication() {
	x.xxx_hidden_Authentication = nil
}

type ApiServiceDescriptor_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// (Required) The URL to connect to.
	Url string
	// (Optional) The method for authentication. If unspecified, access to the
	// url is already being performed in a trusted context (e.g. localhost,
	// private network).
	Authentication *AuthenticationSpec
}

func (b0 ApiServiceDescriptor_builder) Build() *ApiServiceDescriptor {
	m0 := &ApiServiceDescriptor{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Url = b.Url
	x.xxx_hidden_Authentication = b.Authentication
	return m0
}

type AuthenticationSpec struct {
	state              protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Urn     string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	xxx_hidden_Payload []byte                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *AuthenticationSpec) Reset() {
	*x = AuthenticationSpec{}
	mi := &file_org_apache_beam_model_pipeline_v1_endpoints_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationSpec) ProtoMessage() {}

func (x *AuthenticationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_pipeline_v1_endpoints_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *AuthenticationSpec) GetUrn() string {
	if x != nil {
		return x.xxx_hidden_Urn
	}
	return ""
}

func (x *AuthenticationSpec) GetPayload() []byte {
	if x != nil {
		return x.xxx_hidden_Payload
	}
	return nil
}

func (x *AuthenticationSpec) SetUrn(v string) {
	x.xxx_hidden_Urn = v
}

func (x *AuthenticationSpec) SetPayload(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Payload = v
}

type AuthenticationSpec_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// (Required) A URN that describes the accompanying payload.
	// For any URN that is not recognized (by whomever is inspecting
	// it) the parameter payload should be treated as opaque and
	// passed as-is.
	Urn string
	// (Optional) The data specifying any parameters to the URN. If
	// the URN does not require any arguments, this may be omitted.
	Payload []byte
}

func (b0 AuthenticationSpec_builder) Build() *AuthenticationSpec {
	m0 := &AuthenticationSpec{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Urn = b.Urn
	x.xxx_hidden_Payload = b.Payload
	return m0
}

var File_org_apache_beam_model_pipeline_v1_endpoints_proto protoreflect.FileDescriptor

var file_org_apache_beam_model_pipeline_v1_endpoints_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61,
	0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x87, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x69, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x5d, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x40, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x78, 0x0a, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x64, 0x6b, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x31,
	0x3b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_org_apache_beam_model_pipeline_v1_endpoints_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_org_apache_beam_model_pipeline_v1_endpoints_proto_goTypes = []any{
	(*ApiServiceDescriptor)(nil), // 0: org.apache.beam.model.pipeline.v1.ApiServiceDescriptor
	(*AuthenticationSpec)(nil),   // 1: org.apache.beam.model.pipeline.v1.AuthenticationSpec
}
var file_org_apache_beam_model_pipeline_v1_endpoints_proto_depIdxs = []int32{
	1, // 0: org.apache.beam.model.pipeline.v1.ApiServiceDescriptor.authentication:type_name -> org.apache.beam.model.pipeline.v1.AuthenticationSpec
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_org_apache_beam_model_pipeline_v1_endpoints_proto_init() }
func file_org_apache_beam_model_pipeline_v1_endpoints_proto_init() {
	if File_org_apache_beam_model_pipeline_v1_endpoints_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_org_apache_beam_model_pipeline_v1_endpoints_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_org_apache_beam_model_pipeline_v1_endpoints_proto_goTypes,
		DependencyIndexes: file_org_apache_beam_model_pipeline_v1_endpoints_proto_depIdxs,
		MessageInfos:      file_org_apache_beam_model_pipeline_v1_endpoints_proto_msgTypes,
	}.Build()
	File_org_apache_beam_model_pipeline_v1_endpoints_proto = out.File
	file_org_apache_beam_model_pipeline_v1_endpoints_proto_rawDesc = nil
	file_org_apache_beam_model_pipeline_v1_endpoints_proto_goTypes = nil
	file_org_apache_beam_model_pipeline_v1_endpoints_proto_depIdxs = nil
}
