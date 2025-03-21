//
// Copyright 2023 The BFIA Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: ptp.proto

package ptp

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

type TransportOutbound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"` // 二进制报文
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`       // 状态码
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"` // 状态说明
}

func (x *TransportOutbound) Reset() {
	*x = TransportOutbound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportOutbound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportOutbound) ProtoMessage() {}

func (x *TransportOutbound) ProtoReflect() protoreflect.Message {
	mi := &file_ptp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportOutbound.ProtoReflect.Descriptor instead.
func (*TransportOutbound) Descriptor() ([]byte, []int) {
	return file_ptp_proto_rawDescGZIP(), []int{0}
}

func (x *TransportOutbound) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TransportOutbound) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TransportOutbound) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_ptp_proto protoreflect.FileDescriptor

var file_ptp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x70, 0x63, 0x2e, 0x70, 0x74, 0x70, 0x22, 0x5b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2d, 0x69, 0x6e, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ptp_proto_rawDescOnce sync.Once
	file_ptp_proto_rawDescData = file_ptp_proto_rawDesc
)

func file_ptp_proto_rawDescGZIP() []byte {
	file_ptp_proto_rawDescOnce.Do(func() {
		file_ptp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ptp_proto_rawDescData)
	})
	return file_ptp_proto_rawDescData
}

var file_ptp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ptp_proto_goTypes = []interface{}{
	(*TransportOutbound)(nil), // 0: org.ppc.ptp.TransportOutbound
}
var file_ptp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ptp_proto_init() }
func file_ptp_proto_init() {
	if File_ptp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ptp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportOutbound); i {
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
			RawDescriptor: file_ptp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ptp_proto_goTypes,
		DependencyIndexes: file_ptp_proto_depIdxs,
		MessageInfos:      file_ptp_proto_msgTypes,
	}.Build()
	File_ptp_proto = out.File
	file_ptp_proto_rawDesc = nil
	file_ptp_proto_goTypes = nil
	file_ptp_proto_depIdxs = nil
}
