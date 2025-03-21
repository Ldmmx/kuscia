// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: kuscia/proto/api/v1/interconn/common.proto

package interconn

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CommonResponse defines the response for request.
type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *structpb.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CommonResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

// ComponentIO defines input and output of component.
type ComponentIO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ComponentIO) Reset() {
	*x = ComponentIO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentIO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentIO) ProtoMessage() {}

func (x *ComponentIO) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentIO.ProtoReflect.Descriptor instead.
func (*ComponentIO) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP(), []int{1}
}

func (x *ComponentIO) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ComponentIO) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Config defines parameter configuration for each component, etc.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role       *ConfigRole      `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Initiator  *ConfigInitiator `protobuf:"bytes,2,opt,name=initiator,proto3" json:"initiator,omitempty"`
	JobParams  *ConfigParams    `protobuf:"bytes,3,opt,name=job_params,json=jobParams,proto3" json:"job_params,omitempty"`
	TaskParams *ConfigParams    `protobuf:"bytes,4,opt,name=task_params,json=taskParams,proto3" json:"task_params,omitempty"`
	Version    string           `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetRole() *ConfigRole {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *Config) GetInitiator() *ConfigInitiator {
	if x != nil {
		return x.Initiator
	}
	return nil
}

func (x *Config) GetJobParams() *ConfigParams {
	if x != nil {
		return x.JobParams
	}
	return nil
}

func (x *Config) GetTaskParams() *ConfigParams {
	if x != nil {
		return x.TaskParams
	}
	return nil
}

func (x *Config) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// ConfigInitiator defines config initiator information.
type ConfigInitiator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role   string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	NodeId string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *ConfigInitiator) Reset() {
	*x = ConfigInitiator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigInitiator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigInitiator) ProtoMessage() {}

func (x *ConfigInitiator) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigInitiator.ProtoReflect.Descriptor instead.
func (*ConfigInitiator) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigInitiator) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ConfigInitiator) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

// ConfigInitiator defines config role information.
type ConfigRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arbiter []string `protobuf:"bytes,1,rep,name=arbiter,proto3" json:"arbiter,omitempty"`
	Host    []string `protobuf:"bytes,2,rep,name=host,proto3" json:"host,omitempty"`
	Guest   []string `protobuf:"bytes,3,rep,name=guest,proto3" json:"guest,omitempty"`
}

func (x *ConfigRole) Reset() {
	*x = ConfigRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRole) ProtoMessage() {}

func (x *ConfigRole) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRole.ProtoReflect.Descriptor instead.
func (*ConfigRole) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigRole) GetArbiter() []string {
	if x != nil {
		return x.Arbiter
	}
	return nil
}

func (x *ConfigRole) GetHost() []string {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ConfigRole) GetGuest() []string {
	if x != nil {
		return x.Guest
	}
	return nil
}

// ConfigInitiator defines config params information.
type ConfigParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host    *structpb.Struct `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Arbiter *structpb.Struct `protobuf:"bytes,2,opt,name=arbiter,proto3" json:"arbiter,omitempty"`
	Guest   *structpb.Struct `protobuf:"bytes,3,opt,name=guest,proto3" json:"guest,omitempty"`
	Common  *structpb.Struct `protobuf:"bytes,4,opt,name=common,proto3" json:"common,omitempty"`
}

func (x *ConfigParams) Reset() {
	*x = ConfigParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigParams) ProtoMessage() {}

func (x *ConfigParams) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigParams.ProtoReflect.Descriptor instead.
func (*ConfigParams) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigParams) GetHost() *structpb.Struct {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ConfigParams) GetArbiter() *structpb.Struct {
	if x != nil {
		return x.Arbiter
	}
	return nil
}

func (x *ConfigParams) GetGuest() *structpb.Struct {
	if x != nil {
		return x.Guest
	}
	return nil
}

func (x *ConfigParams) GetCommon() *structpb.Struct {
	if x != nil {
		return x.Common
	}
	return nil
}

// ConfigInitiator defines config Resources information.
type ConfigResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disk   int32 `protobuf:"varint,1,opt,name=disk,proto3" json:"disk,omitempty"`
	Memory int32 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Cpu    int32 `protobuf:"varint,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
}

func (x *ConfigResources) Reset() {
	*x = ConfigResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResources) ProtoMessage() {}

func (x *ConfigResources) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResources.ProtoReflect.Descriptor instead.
func (*ConfigResources) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigResources) GetDisk() int32 {
	if x != nil {
		return x.Disk
	}
	return 0
}

func (x *ConfigResources) GetMemory() int32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *ConfigResources) GetCpu() int32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

var File_kuscia_proto_api_v1_interconn_common_proto protoreflect.FileDescriptor

var file_kuscia_proto_api_v1_interconn_common_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6b, 0x75,
	0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x4f, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0xc9, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6b,
	0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x4c, 0x0a,
	0x09, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0a, 0x6a,
	0x6f, 0x62, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x09, 0x6a, 0x6f,
	0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6b,
	0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x3e, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22,
	0x50, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xce, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x07, 0x61, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x61, 0x72, 0x62, 0x69, 0x74,
	0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x67, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x63, 0x70, 0x75, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6b, 0x75, 0x73,
	0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kuscia_proto_api_v1_interconn_common_proto_rawDescOnce sync.Once
	file_kuscia_proto_api_v1_interconn_common_proto_rawDescData = file_kuscia_proto_api_v1_interconn_common_proto_rawDesc
)

func file_kuscia_proto_api_v1_interconn_common_proto_rawDescGZIP() []byte {
	file_kuscia_proto_api_v1_interconn_common_proto_rawDescOnce.Do(func() {
		file_kuscia_proto_api_v1_interconn_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuscia_proto_api_v1_interconn_common_proto_rawDescData)
	})
	return file_kuscia_proto_api_v1_interconn_common_proto_rawDescData
}

var file_kuscia_proto_api_v1_interconn_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_kuscia_proto_api_v1_interconn_common_proto_goTypes = []interface{}{
	(*CommonResponse)(nil),  // 0: kuscia.proto.api.v1.interconn.CommonResponse
	(*ComponentIO)(nil),     // 1: kuscia.proto.api.v1.interconn.ComponentIO
	(*Config)(nil),          // 2: kuscia.proto.api.v1.interconn.Config
	(*ConfigInitiator)(nil), // 3: kuscia.proto.api.v1.interconn.ConfigInitiator
	(*ConfigRole)(nil),      // 4: kuscia.proto.api.v1.interconn.ConfigRole
	(*ConfigParams)(nil),    // 5: kuscia.proto.api.v1.interconn.ConfigParams
	(*ConfigResources)(nil), // 6: kuscia.proto.api.v1.interconn.ConfigResources
	(*structpb.Struct)(nil), // 7: google.protobuf.Struct
}
var file_kuscia_proto_api_v1_interconn_common_proto_depIdxs = []int32{
	7, // 0: kuscia.proto.api.v1.interconn.CommonResponse.data:type_name -> google.protobuf.Struct
	4, // 1: kuscia.proto.api.v1.interconn.Config.role:type_name -> kuscia.proto.api.v1.interconn.ConfigRole
	3, // 2: kuscia.proto.api.v1.interconn.Config.initiator:type_name -> kuscia.proto.api.v1.interconn.ConfigInitiator
	5, // 3: kuscia.proto.api.v1.interconn.Config.job_params:type_name -> kuscia.proto.api.v1.interconn.ConfigParams
	5, // 4: kuscia.proto.api.v1.interconn.Config.task_params:type_name -> kuscia.proto.api.v1.interconn.ConfigParams
	7, // 5: kuscia.proto.api.v1.interconn.ConfigParams.host:type_name -> google.protobuf.Struct
	7, // 6: kuscia.proto.api.v1.interconn.ConfigParams.arbiter:type_name -> google.protobuf.Struct
	7, // 7: kuscia.proto.api.v1.interconn.ConfigParams.guest:type_name -> google.protobuf.Struct
	7, // 8: kuscia.proto.api.v1.interconn.ConfigParams.common:type_name -> google.protobuf.Struct
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_kuscia_proto_api_v1_interconn_common_proto_init() }
func file_kuscia_proto_api_v1_interconn_common_proto_init() {
	if File_kuscia_proto_api_v1_interconn_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentIO); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigInitiator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRole); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kuscia_proto_api_v1_interconn_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResources); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kuscia_proto_api_v1_interconn_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kuscia_proto_api_v1_interconn_common_proto_goTypes,
		DependencyIndexes: file_kuscia_proto_api_v1_interconn_common_proto_depIdxs,
		MessageInfos:      file_kuscia_proto_api_v1_interconn_common_proto_msgTypes,
	}.Build()
	File_kuscia_proto_api_v1_interconn_common_proto = out.File
	file_kuscia_proto_api_v1_interconn_common_proto_rawDesc = nil
	file_kuscia_proto_api_v1_interconn_common_proto_goTypes = nil
	file_kuscia_proto_api_v1_interconn_common_proto_depIdxs = nil
}
