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
// Protocol Buffers describing the Expansion API, an api for expanding
// transforms in a remote SDK.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.27.3
// source: org/apache/beam/model/job_management/v1/beam_expansion_api.proto

//go:build protoopaque

package jobmanagement_v1

import (
	pipeline_v1 "github.com/apache/beam/sdks/v2/go/pkg/beam/model/pipeline_v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExpansionRequest struct {
	state                          protoimpl.MessageState  `protogen:"opaque.v1"`
	xxx_hidden_Components          *pipeline_v1.Components `protobuf:"bytes,1,opt,name=components,proto3" json:"components,omitempty"`
	xxx_hidden_Transform           *pipeline_v1.PTransform `protobuf:"bytes,2,opt,name=transform,proto3" json:"transform,omitempty"`
	xxx_hidden_Namespace           string                  `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	xxx_hidden_OutputCoderRequests map[string]string       `protobuf:"bytes,4,rep,name=output_coder_requests,json=outputCoderRequests,proto3" json:"output_coder_requests,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_Requirements        []string                `protobuf:"bytes,5,rep,name=requirements,proto3" json:"requirements,omitempty"`
	xxx_hidden_PipelineOptions     *structpb.Struct        `protobuf:"bytes,6,opt,name=pipeline_options,json=pipelineOptions,proto3" json:"pipeline_options,omitempty"`
	unknownFields                  protoimpl.UnknownFields
	sizeCache                      protoimpl.SizeCache
}

func (x *ExpansionRequest) Reset() {
	*x = ExpansionRequest{}
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpansionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpansionRequest) ProtoMessage() {}

func (x *ExpansionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ExpansionRequest) GetComponents() *pipeline_v1.Components {
	if x != nil {
		return x.xxx_hidden_Components
	}
	return nil
}

func (x *ExpansionRequest) GetTransform() *pipeline_v1.PTransform {
	if x != nil {
		return x.xxx_hidden_Transform
	}
	return nil
}

func (x *ExpansionRequest) GetNamespace() string {
	if x != nil {
		return x.xxx_hidden_Namespace
	}
	return ""
}

func (x *ExpansionRequest) GetOutputCoderRequests() map[string]string {
	if x != nil {
		return x.xxx_hidden_OutputCoderRequests
	}
	return nil
}

func (x *ExpansionRequest) GetRequirements() []string {
	if x != nil {
		return x.xxx_hidden_Requirements
	}
	return nil
}

func (x *ExpansionRequest) GetPipelineOptions() *structpb.Struct {
	if x != nil {
		return x.xxx_hidden_PipelineOptions
	}
	return nil
}

func (x *ExpansionRequest) SetComponents(v *pipeline_v1.Components) {
	x.xxx_hidden_Components = v
}

func (x *ExpansionRequest) SetTransform(v *pipeline_v1.PTransform) {
	x.xxx_hidden_Transform = v
}

func (x *ExpansionRequest) SetNamespace(v string) {
	x.xxx_hidden_Namespace = v
}

func (x *ExpansionRequest) SetOutputCoderRequests(v map[string]string) {
	x.xxx_hidden_OutputCoderRequests = v
}

func (x *ExpansionRequest) SetRequirements(v []string) {
	x.xxx_hidden_Requirements = v
}

func (x *ExpansionRequest) SetPipelineOptions(v *structpb.Struct) {
	x.xxx_hidden_PipelineOptions = v
}

func (x *ExpansionRequest) HasComponents() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Components != nil
}

func (x *ExpansionRequest) HasTransform() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Transform != nil
}

func (x *ExpansionRequest) HasPipelineOptions() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_PipelineOptions != nil
}

func (x *ExpansionRequest) ClearComponents() {
	x.xxx_hidden_Components = nil
}

func (x *ExpansionRequest) ClearTransform() {
	x.xxx_hidden_Transform = nil
}

func (x *ExpansionRequest) ClearPipelineOptions() {
	x.xxx_hidden_PipelineOptions = nil
}

type ExpansionRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Set of components needed to interpret the transform, or which
	// may be useful for its expansion.  This includes the input
	// PCollections (if any) to the to-be-expanded transform, along
	// with their coders and windowing strategies.
	Components *pipeline_v1.Components
	// The actual PTransform to be expaneded according to its spec.
	// Its input should be set, but its subtransforms and outputs
	// should not be.
	Transform *pipeline_v1.PTransform
	// A namespace (prefix) to use for the id of any newly created
	// components.
	Namespace string
	// (Optional) Map from a local output tag to a coder id.
	// If it is set, asks the expansion service to use the given
	// coders for the output PCollections. Note that the request
	// may not be fulfilled.
	OutputCoderRequests map[string]string
	// A set of requirements that must be used by the expansion service to
	// interpret the components provided with this request.
	Requirements []string
	// (Optional) A set of Pipeline Options that should be used
	// when expanding this transform.
	PipelineOptions *structpb.Struct
}

func (b0 ExpansionRequest_builder) Build() *ExpansionRequest {
	m0 := &ExpansionRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Components = b.Components
	x.xxx_hidden_Transform = b.Transform
	x.xxx_hidden_Namespace = b.Namespace
	x.xxx_hidden_OutputCoderRequests = b.OutputCoderRequests
	x.xxx_hidden_Requirements = b.Requirements
	x.xxx_hidden_PipelineOptions = b.PipelineOptions
	return m0
}

type ExpansionResponse struct {
	state                   protoimpl.MessageState  `protogen:"opaque.v1"`
	xxx_hidden_Components   *pipeline_v1.Components `protobuf:"bytes,1,opt,name=components,proto3" json:"components,omitempty"`
	xxx_hidden_Transform    *pipeline_v1.PTransform `protobuf:"bytes,2,opt,name=transform,proto3" json:"transform,omitempty"`
	xxx_hidden_Requirements []string                `protobuf:"bytes,3,rep,name=requirements,proto3" json:"requirements,omitempty"`
	xxx_hidden_Error        string                  `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *ExpansionResponse) Reset() {
	*x = ExpansionResponse{}
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpansionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpansionResponse) ProtoMessage() {}

func (x *ExpansionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ExpansionResponse) GetComponents() *pipeline_v1.Components {
	if x != nil {
		return x.xxx_hidden_Components
	}
	return nil
}

func (x *ExpansionResponse) GetTransform() *pipeline_v1.PTransform {
	if x != nil {
		return x.xxx_hidden_Transform
	}
	return nil
}

func (x *ExpansionResponse) GetRequirements() []string {
	if x != nil {
		return x.xxx_hidden_Requirements
	}
	return nil
}

func (x *ExpansionResponse) GetError() string {
	if x != nil {
		return x.xxx_hidden_Error
	}
	return ""
}

func (x *ExpansionResponse) SetComponents(v *pipeline_v1.Components) {
	x.xxx_hidden_Components = v
}

func (x *ExpansionResponse) SetTransform(v *pipeline_v1.PTransform) {
	x.xxx_hidden_Transform = v
}

func (x *ExpansionResponse) SetRequirements(v []string) {
	x.xxx_hidden_Requirements = v
}

func (x *ExpansionResponse) SetError(v string) {
	x.xxx_hidden_Error = v
}

func (x *ExpansionResponse) HasComponents() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Components != nil
}

func (x *ExpansionResponse) HasTransform() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Transform != nil
}

func (x *ExpansionResponse) ClearComponents() {
	x.xxx_hidden_Components = nil
}

func (x *ExpansionResponse) ClearTransform() {
	x.xxx_hidden_Transform = nil
}

type ExpansionResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Set of components needed to execute the expanded transform,
	// including the (original) inputs, outputs, and subtransforms.
	Components *pipeline_v1.Components
	// The expanded transform itself, with references to its outputs
	// and subtransforms.
	Transform *pipeline_v1.PTransform
	// A set of requirements that must be appended to this pipeline's
	// requirements.
	Requirements []string
	// (Optional) An string representation of any error encountered while
	// attempting to expand this transform.
	Error string
}

func (b0 ExpansionResponse_builder) Build() *ExpansionResponse {
	m0 := &ExpansionResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Components = b.Components
	x.xxx_hidden_Transform = b.Transform
	x.xxx_hidden_Requirements = b.Requirements
	x.xxx_hidden_Error = b.Error
	return m0
}

type DiscoverSchemaTransformRequest struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DiscoverSchemaTransformRequest) Reset() {
	*x = DiscoverSchemaTransformRequest{}
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverSchemaTransformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverSchemaTransformRequest) ProtoMessage() {}

func (x *DiscoverSchemaTransformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type DiscoverSchemaTransformRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 DiscoverSchemaTransformRequest_builder) Build() *DiscoverSchemaTransformRequest {
	m0 := &DiscoverSchemaTransformRequest{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type SchemaTransformConfig struct {
	state                             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ConfigSchema           *pipeline_v1.Schema    `protobuf:"bytes,1,opt,name=config_schema,json=configSchema,proto3" json:"config_schema,omitempty"`
	xxx_hidden_InputPcollectionNames  []string               `protobuf:"bytes,2,rep,name=input_pcollection_names,json=inputPcollectionNames,proto3" json:"input_pcollection_names,omitempty"`
	xxx_hidden_OutputPcollectionNames []string               `protobuf:"bytes,3,rep,name=output_pcollection_names,json=outputPcollectionNames,proto3" json:"output_pcollection_names,omitempty"`
	xxx_hidden_Description            string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields                     protoimpl.UnknownFields
	sizeCache                         protoimpl.SizeCache
}

func (x *SchemaTransformConfig) Reset() {
	*x = SchemaTransformConfig{}
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchemaTransformConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaTransformConfig) ProtoMessage() {}

func (x *SchemaTransformConfig) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SchemaTransformConfig) GetConfigSchema() *pipeline_v1.Schema {
	if x != nil {
		return x.xxx_hidden_ConfigSchema
	}
	return nil
}

func (x *SchemaTransformConfig) GetInputPcollectionNames() []string {
	if x != nil {
		return x.xxx_hidden_InputPcollectionNames
	}
	return nil
}

func (x *SchemaTransformConfig) GetOutputPcollectionNames() []string {
	if x != nil {
		return x.xxx_hidden_OutputPcollectionNames
	}
	return nil
}

func (x *SchemaTransformConfig) GetDescription() string {
	if x != nil {
		return x.xxx_hidden_Description
	}
	return ""
}

func (x *SchemaTransformConfig) SetConfigSchema(v *pipeline_v1.Schema) {
	x.xxx_hidden_ConfigSchema = v
}

func (x *SchemaTransformConfig) SetInputPcollectionNames(v []string) {
	x.xxx_hidden_InputPcollectionNames = v
}

func (x *SchemaTransformConfig) SetOutputPcollectionNames(v []string) {
	x.xxx_hidden_OutputPcollectionNames = v
}

func (x *SchemaTransformConfig) SetDescription(v string) {
	x.xxx_hidden_Description = v
}

func (x *SchemaTransformConfig) HasConfigSchema() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_ConfigSchema != nil
}

func (x *SchemaTransformConfig) ClearConfigSchema() {
	x.xxx_hidden_ConfigSchema = nil
}

type SchemaTransformConfig_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Config schema of the SchemaTransform
	ConfigSchema *pipeline_v1.Schema
	// Names of input PCollections
	InputPcollectionNames []string
	// Names of output PCollections
	OutputPcollectionNames []string
	// Description of this transform and usage used for documentation.
	// May be markdown formatted.
	// Note that configuration parameters may also have documentation attached
	// as part of the config_schema.
	Description string
}

func (b0 SchemaTransformConfig_builder) Build() *SchemaTransformConfig {
	m0 := &SchemaTransformConfig{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_ConfigSchema = b.ConfigSchema
	x.xxx_hidden_InputPcollectionNames = b.InputPcollectionNames
	x.xxx_hidden_OutputPcollectionNames = b.OutputPcollectionNames
	x.xxx_hidden_Description = b.Description
	return m0
}

type DiscoverSchemaTransformResponse struct {
	state                             protoimpl.MessageState            `protogen:"opaque.v1"`
	xxx_hidden_SchemaTransformConfigs map[string]*SchemaTransformConfig `protobuf:"bytes,1,rep,name=schema_transform_configs,json=schemaTransformConfigs,proto3" json:"schema_transform_configs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_Error                  string                            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields                     protoimpl.UnknownFields
	sizeCache                         protoimpl.SizeCache
}

func (x *DiscoverSchemaTransformResponse) Reset() {
	*x = DiscoverSchemaTransformResponse{}
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverSchemaTransformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverSchemaTransformResponse) ProtoMessage() {}

func (x *DiscoverSchemaTransformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DiscoverSchemaTransformResponse) GetSchemaTransformConfigs() map[string]*SchemaTransformConfig {
	if x != nil {
		return x.xxx_hidden_SchemaTransformConfigs
	}
	return nil
}

func (x *DiscoverSchemaTransformResponse) GetError() string {
	if x != nil {
		return x.xxx_hidden_Error
	}
	return ""
}

func (x *DiscoverSchemaTransformResponse) SetSchemaTransformConfigs(v map[string]*SchemaTransformConfig) {
	x.xxx_hidden_SchemaTransformConfigs = v
}

func (x *DiscoverSchemaTransformResponse) SetError(v string) {
	x.xxx_hidden_Error = v
}

type DiscoverSchemaTransformResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// A mapping from SchemaTransform ID to schema transform config of discovered
	// SchemaTransforms
	SchemaTransformConfigs map[string]*SchemaTransformConfig
	// If list of identifies are empty, this may contain an error.
	Error string
}

func (b0 DiscoverSchemaTransformResponse_builder) Build() *DiscoverSchemaTransformResponse {
	m0 := &DiscoverSchemaTransformResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_SchemaTransformConfigs = b.SchemaTransformConfigs
	x.xxx_hidden_Error = b.Error
	return m0
}

var File_org_apache_beam_model_job_management_v1_beam_expansion_api_proto protoreflect.FileDescriptor

var file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_rawDesc = []byte{
	0x0a, 0x40, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61,
	0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x5f, 0x65,
	0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62,
	0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x04,
	0x0a, 0x10, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x4b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a,
	0x15, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x10, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x46, 0x0a, 0x18, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xe9, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x20, 0x0a, 0x1e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfb,
	0x01, 0x0a, 0x15, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61,
	0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x70, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x50, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x38, 0x0a, 0x18, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x16, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xda, 0x02, 0x0a,
	0x1f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x99, 0x01, 0x0a, 0x18, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x5f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x61,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x1a, 0x84, 0x01, 0x0a, 0x1b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x4f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x61,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xae, 0x02, 0x0a, 0x10, 0x45, 0x78,
	0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75,
	0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x34, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa2, 0x01, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x42, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62,
	0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78,
	0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x86, 0x01, 0x0a, 0x22, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x64, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x76, 0x31, 0x3b, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_goTypes = []any{
	(*ExpansionRequest)(nil),                // 0: org.apache.beam.model.expansion.v1.ExpansionRequest
	(*ExpansionResponse)(nil),               // 1: org.apache.beam.model.expansion.v1.ExpansionResponse
	(*DiscoverSchemaTransformRequest)(nil),  // 2: org.apache.beam.model.expansion.v1.DiscoverSchemaTransformRequest
	(*SchemaTransformConfig)(nil),           // 3: org.apache.beam.model.expansion.v1.SchemaTransformConfig
	(*DiscoverSchemaTransformResponse)(nil), // 4: org.apache.beam.model.expansion.v1.DiscoverSchemaTransformResponse
	nil,                                     // 5: org.apache.beam.model.expansion.v1.ExpansionRequest.OutputCoderRequestsEntry
	nil,                                     // 6: org.apache.beam.model.expansion.v1.DiscoverSchemaTransformResponse.SchemaTransformConfigsEntry
	(*pipeline_v1.Components)(nil),          // 7: org.apache.beam.model.pipeline.v1.Components
	(*pipeline_v1.PTransform)(nil),          // 8: org.apache.beam.model.pipeline.v1.PTransform
	(*structpb.Struct)(nil),                 // 9: google.protobuf.Struct
	(*pipeline_v1.Schema)(nil),              // 10: org.apache.beam.model.pipeline.v1.Schema
}
var file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_depIdxs = []int32{
	7,  // 0: org.apache.beam.model.expansion.v1.ExpansionRequest.components:type_name -> org.apache.beam.model.pipeline.v1.Components
	8,  // 1: org.apache.beam.model.expansion.v1.ExpansionRequest.transform:type_name -> org.apache.beam.model.pipeline.v1.PTransform
	5,  // 2: org.apache.beam.model.expansion.v1.ExpansionRequest.output_coder_requests:type_name -> org.apache.beam.model.expansion.v1.ExpansionRequest.OutputCoderRequestsEntry
	9,  // 3: org.apache.beam.model.expansion.v1.ExpansionRequest.pipeline_options:type_name -> google.protobuf.Struct
	7,  // 4: org.apache.beam.model.expansion.v1.ExpansionResponse.components:type_name -> org.apache.beam.model.pipeline.v1.Components
	8,  // 5: org.apache.beam.model.expansion.v1.ExpansionResponse.transform:type_name -> org.apache.beam.model.pipeline.v1.PTransform
	10, // 6: org.apache.beam.model.expansion.v1.SchemaTransformConfig.config_schema:type_name -> org.apache.beam.model.pipeline.v1.Schema
	6,  // 7: org.apache.beam.model.expansion.v1.DiscoverSchemaTransformResponse.schema_transform_configs:type_name -> org.apache.beam.model.expansion.v1.DiscoverSchemaTransformResponse.SchemaTransformConfigsEntry
	3,  // 8: org.apache.beam.model.expansion.v1.DiscoverSchemaTransformResponse.SchemaTransformConfigsEntry.value:type_name -> org.apache.beam.model.expansion.v1.SchemaTransformConfig
	0,  // 9: org.apache.beam.model.expansion.v1.ExpansionService.Expand:input_type -> org.apache.beam.model.expansion.v1.ExpansionRequest
	2,  // 10: org.apache.beam.model.expansion.v1.ExpansionService.DiscoverSchemaTransform:input_type -> org.apache.beam.model.expansion.v1.DiscoverSchemaTransformRequest
	1,  // 11: org.apache.beam.model.expansion.v1.ExpansionService.Expand:output_type -> org.apache.beam.model.expansion.v1.ExpansionResponse
	4,  // 12: org.apache.beam.model.expansion.v1.ExpansionService.DiscoverSchemaTransform:output_type -> org.apache.beam.model.expansion.v1.DiscoverSchemaTransformResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_init() }
func file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_init() {
	if File_org_apache_beam_model_job_management_v1_beam_expansion_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_goTypes,
		DependencyIndexes: file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_depIdxs,
		MessageInfos:      file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_msgTypes,
	}.Build()
	File_org_apache_beam_model_job_management_v1_beam_expansion_api_proto = out.File
	file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_rawDesc = nil
	file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_goTypes = nil
	file_org_apache_beam_model_job_management_v1_beam_expansion_api_proto_depIdxs = nil
}
