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
// source: kuscia/proto/api/v1alpha1/appconfig/app_config.proto

package appconfig

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Service represents the service address corresponding to the port.
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of port.
	PortName string `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	// Endpoint list corresponding to the port.
	Endpoints []string `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetPortName() string {
	if x != nil {
		return x.PortName
	}
	return ""
}

func (x *Service) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

// Party represents the basic information of the party.
type Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of party.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// role carried by party. Examples: client, server...
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// List of services exposed by pod.
	Services []*Service `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *Party) Reset() {
	*x = Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Party) ProtoMessage() {}

func (x *Party) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Party.ProtoReflect.Descriptor instead.
func (*Party) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescGZIP(), []int{1}
}

func (x *Party) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Party) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Party) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

// ClusterDefine represents the information of all parties.
type ClusterDefine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Basic information of all parties.
	Parties []*Party `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
	// index of self party.
	SelfPartyIdx int32 `protobuf:"varint,2,opt,name=self_party_idx,json=selfPartyIdx,proto3" json:"self_party_idx,omitempty"`
	// index of self endpoint.
	SelfEndpointIdx int32 `protobuf:"varint,3,opt,name=self_endpoint_idx,json=selfEndpointIdx,proto3" json:"self_endpoint_idx,omitempty"`
}

func (x *ClusterDefine) Reset() {
	*x = ClusterDefine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDefine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDefine) ProtoMessage() {}

func (x *ClusterDefine) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDefine.ProtoReflect.Descriptor instead.
func (*ClusterDefine) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterDefine) GetParties() []*Party {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *ClusterDefine) GetSelfPartyIdx() int32 {
	if x != nil {
		return x.SelfPartyIdx
	}
	return 0
}

func (x *ClusterDefine) GetSelfEndpointIdx() int32 {
	if x != nil {
		return x.SelfEndpointIdx
	}
	return 0
}

// Port represents an allocated port for pod.
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each named port in a pod must have a unique name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Number of port allocated for pod.
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Scope of port. Must be Cluster,Domain,Local.
	// Defaults to "Local".
	// +optional
	Scope string `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	// Protocol for port. Must be HTTP,GRPC.
	// Defaults to "HTTP".
	// +optional
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescGZIP(), []int{3}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Port) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Port) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

// AllocatedPorts represents allocated ports for pod.
type AllocatedPorts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Allocated ports.
	Ports []*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *AllocatedPorts) Reset() {
	*x = AllocatedPorts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocatedPorts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocatedPorts) ProtoMessage() {}

func (x *AllocatedPorts) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocatedPorts.ProtoReflect.Descriptor instead.
func (*AllocatedPorts) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescGZIP(), []int{4}
}

func (x *AllocatedPorts) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

var File_kuscia_proto_api_v1alpha1_appconfig_app_config_proto protoreflect.FileDescriptor

var file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x44, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x22, 0x79, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xa7, 0x01, 0x0a,
	0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x44,
	0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65,
	0x6c, 0x66, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x78, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65,
	0x6c, 0x66, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x65, 0x6c, 0x66, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x78, 0x22, 0x60, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x51, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x05, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x75, 0x73, 0x63,
	0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x42, 0x5e, 0x0a, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescOnce sync.Once
	file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescData = file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDesc
)

func file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescGZIP() []byte {
	file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescOnce.Do(func() {
		file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescData)
	})
	return file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDescData
}

var file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_goTypes = []interface{}{
	(*Service)(nil),        // 0: kuscia.proto.api.v1alpha1.appconfig.Service
	(*Party)(nil),          // 1: kuscia.proto.api.v1alpha1.appconfig.Party
	(*ClusterDefine)(nil),  // 2: kuscia.proto.api.v1alpha1.appconfig.ClusterDefine
	(*Port)(nil),           // 3: kuscia.proto.api.v1alpha1.appconfig.Port
	(*AllocatedPorts)(nil), // 4: kuscia.proto.api.v1alpha1.appconfig.AllocatedPorts
}
var file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_depIdxs = []int32{
	0, // 0: kuscia.proto.api.v1alpha1.appconfig.Party.services:type_name -> kuscia.proto.api.v1alpha1.appconfig.Service
	1, // 1: kuscia.proto.api.v1alpha1.appconfig.ClusterDefine.parties:type_name -> kuscia.proto.api.v1alpha1.appconfig.Party
	3, // 2: kuscia.proto.api.v1alpha1.appconfig.AllocatedPorts.ports:type_name -> kuscia.proto.api.v1alpha1.appconfig.Port
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_init() }
func file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_init() {
	if File_kuscia_proto_api_v1alpha1_appconfig_app_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Party); i {
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
		file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDefine); i {
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
		file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocatedPorts); i {
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
			RawDescriptor: file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_goTypes,
		DependencyIndexes: file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_depIdxs,
		MessageInfos:      file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_msgTypes,
	}.Build()
	File_kuscia_proto_api_v1alpha1_appconfig_app_config_proto = out.File
	file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_rawDesc = nil
	file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_goTypes = nil
	file_kuscia_proto_api_v1alpha1_appconfig_app_config_proto_depIdxs = nil
}
