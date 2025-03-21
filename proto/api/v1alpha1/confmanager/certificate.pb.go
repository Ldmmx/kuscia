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
// source: kuscia/proto/api/v1alpha1/confmanager/certificate.proto

package confmanager

import (
	v1alpha1 "github.com/secretflow/kuscia/proto/api/v1alpha1"
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

type GenerateKeyCertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common Name, required
	CommonName string `protobuf:"bytes,1,opt,name=common_name,json=commonName,proto3" json:"common_name,omitempty"`
	// Country, optional
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	// Organization, optional
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
	// Organization Unit, optional
	OrganizationUnit string `protobuf:"bytes,4,opt,name=organization_unit,json=organizationUnit,proto3" json:"organization_unit,omitempty"`
	// Locality, optional
	Locality string `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	// Province, optional
	Province string `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty"`
	// Street Address, optional
	StreetAddress string `protobuf:"bytes,7,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	// Valid Duration Seconds, optional, from now, default: 1 day
	DurationSec int64 `protobuf:"varint,8,opt,name=duration_sec,json=durationSec,proto3" json:"duration_sec,omitempty"`
	// Key Type, Enum: [PKCS#1, PKCS#8], optional, default: PKCS#1
	KeyType string `protobuf:"bytes,9,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
}

func (x *GenerateKeyCertsRequest) Reset() {
	*x = GenerateKeyCertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeyCertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeyCertsRequest) ProtoMessage() {}

func (x *GenerateKeyCertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeyCertsRequest.ProtoReflect.Descriptor instead.
func (*GenerateKeyCertsRequest) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateKeyCertsRequest) GetCommonName() string {
	if x != nil {
		return x.CommonName
	}
	return ""
}

func (x *GenerateKeyCertsRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GenerateKeyCertsRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *GenerateKeyCertsRequest) GetOrganizationUnit() string {
	if x != nil {
		return x.OrganizationUnit
	}
	return ""
}

func (x *GenerateKeyCertsRequest) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *GenerateKeyCertsRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *GenerateKeyCertsRequest) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *GenerateKeyCertsRequest) GetDurationSec() int64 {
	if x != nil {
		return x.DurationSec
	}
	return 0
}

func (x *GenerateKeyCertsRequest) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

type GenerateKeyCertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *v1alpha1.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The generate private key.Default PKCS#1. Base64 Encoded.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The cert chain of generate cert file.The first is the generate cert, The last is domain root ca cert.Base64 Encoded.
	CertChain []string `protobuf:"bytes,3,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
}

func (x *GenerateKeyCertsResponse) Reset() {
	*x = GenerateKeyCertsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeyCertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeyCertsResponse) ProtoMessage() {}

func (x *GenerateKeyCertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeyCertsResponse.ProtoReflect.Descriptor instead.
func (*GenerateKeyCertsResponse) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateKeyCertsResponse) GetStatus() *v1alpha1.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GenerateKeyCertsResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GenerateKeyCertsResponse) GetCertChain() []string {
	if x != nil {
		return x.CertChain
	}
	return nil
}

var File_kuscia_proto_api_v1alpha1_confmanager_certificate_proto protoreflect.FileDescriptor

var file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x6b, 0x75, 0x73, 0x63, 0x69,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x1a, 0x26, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x86, 0x01,
	0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x43, 0x65, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x75, 0x73,
	0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x72,
	0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x32, 0xaa, 0x01, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x43, 0x65, 0x72,
	0x74, 0x73, 0x12, 0x3e, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x62, 0x0a, 0x23, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescOnce sync.Once
	file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescData = file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDesc
)

func file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescGZIP() []byte {
	file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescOnce.Do(func() {
		file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescData)
	})
	return file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDescData
}

var file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_goTypes = []interface{}{
	(*GenerateKeyCertsRequest)(nil),  // 0: kuscia.proto.api.v1alpha1.confmanager.GenerateKeyCertsRequest
	(*GenerateKeyCertsResponse)(nil), // 1: kuscia.proto.api.v1alpha1.confmanager.GenerateKeyCertsResponse
	(*v1alpha1.Status)(nil),          // 2: kuscia.proto.api.v1alpha1.Status
}
var file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_depIdxs = []int32{
	2, // 0: kuscia.proto.api.v1alpha1.confmanager.GenerateKeyCertsResponse.status:type_name -> kuscia.proto.api.v1alpha1.Status
	0, // 1: kuscia.proto.api.v1alpha1.confmanager.CertificateService.GenerateKeyCerts:input_type -> kuscia.proto.api.v1alpha1.confmanager.GenerateKeyCertsRequest
	1, // 2: kuscia.proto.api.v1alpha1.confmanager.CertificateService.GenerateKeyCerts:output_type -> kuscia.proto.api.v1alpha1.confmanager.GenerateKeyCertsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_init() }
func file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_init() {
	if File_kuscia_proto_api_v1alpha1_confmanager_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeyCertsRequest); i {
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
		file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeyCertsResponse); i {
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
			RawDescriptor: file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_goTypes,
		DependencyIndexes: file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_depIdxs,
		MessageInfos:      file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_msgTypes,
	}.Build()
	File_kuscia_proto_api_v1alpha1_confmanager_certificate_proto = out.File
	file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_rawDesc = nil
	file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_goTypes = nil
	file_kuscia_proto_api_v1alpha1_confmanager_certificate_proto_depIdxs = nil
}
