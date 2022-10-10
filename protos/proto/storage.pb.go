// Copyright Rivtower Technologies LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: storage.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Regions int32

const (
	Regions_GLOBAL                        Regions = 0
	Regions_TRANSACTIONS                  Regions = 1
	Regions_HEADERS                       Regions = 2
	Regions_BODIES                        Regions = 3
	Regions_BLOCK_HASH                    Regions = 4
	Regions_PROOF                         Regions = 5
	Regions_RESULT                        Regions = 6
	Regions_TRANSACTION_HASH2BLOCK_HEIGHT Regions = 7
	Regions_BLOCK_HASH2BLOCK_HEIGHT       Regions = 8 // In SQL db, reuse 4
	Regions_TRANSACTION_INDEX             Regions = 9
	Regions_COMPACT_BLOCK                 Regions = 10
	Regions_FULL_BLOCK                    Regions = 11
	Regions_All_BLOCK_DATA                Regions = 12
	Regions_BUTTON                        Regions = 13
)

// Enum value maps for Regions.
var (
	Regions_name = map[int32]string{
		0:  "GLOBAL",
		1:  "TRANSACTIONS",
		2:  "HEADERS",
		3:  "BODIES",
		4:  "BLOCK_HASH",
		5:  "PROOF",
		6:  "RESULT",
		7:  "TRANSACTION_HASH2BLOCK_HEIGHT",
		8:  "BLOCK_HASH2BLOCK_HEIGHT",
		9:  "TRANSACTION_INDEX",
		10: "COMPACT_BLOCK",
		11: "FULL_BLOCK",
		12: "All_BLOCK_DATA",
		13: "BUTTON",
	}
	Regions_value = map[string]int32{
		"GLOBAL":                        0,
		"TRANSACTIONS":                  1,
		"HEADERS":                       2,
		"BODIES":                        3,
		"BLOCK_HASH":                    4,
		"PROOF":                         5,
		"RESULT":                        6,
		"TRANSACTION_HASH2BLOCK_HEIGHT": 7,
		"BLOCK_HASH2BLOCK_HEIGHT":       8,
		"TRANSACTION_INDEX":             9,
		"COMPACT_BLOCK":                 10,
		"FULL_BLOCK":                    11,
		"All_BLOCK_DATA":                12,
		"BUTTON":                        13,
	}
)

func (x Regions) Enum() *Regions {
	p := new(Regions)
	*p = x
	return p
}

func (x Regions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Regions) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_proto_enumTypes[0].Descriptor()
}

func (Regions) Type() protoreflect.EnumType {
	return &file_storage_proto_enumTypes[0]
}

func (x Regions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Regions.Descriptor instead.
func (Regions) EnumDescriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region uint32 `protobuf:"varint,1,opt,name=region,proto3" json:"region,omitempty"`
	Key    []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetRegion() uint32 {
	if x != nil {
		return x.Region
	}
	return 0
}

func (x *Content) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Content) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ExtKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region uint32 `protobuf:"varint,1,opt,name=region,proto3" json:"region,omitempty"`
	Key    []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ExtKey) Reset() {
	*x = ExtKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtKey) ProtoMessage() {}

func (x *ExtKey) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtKey.ProtoReflect.Descriptor instead.
func (*ExtKey) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *ExtKey) GetRegion() uint32 {
	if x != nil {
		return x.Region
	}
	return 0
}

func (x *ExtKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *StatusCode `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Value  []byte      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Value) GetStatus() *StatusCode {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Value) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x32, 0x0a, 0x06, 0x45, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x49, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2a, 0x81, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0a, 0x0a, 0x06,
	0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45,
	0x41, 0x44, 0x45, 0x52, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4f, 0x44, 0x49, 0x45,
	0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x48, 0x41, 0x53,
	0x48, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x4f, 0x4f, 0x46, 0x10, 0x05, 0x12, 0x0a,
	0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x06, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x32, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x48, 0x45, 0x49, 0x47, 0x48, 0x54, 0x10, 0x07, 0x12, 0x1b, 0x0a,
	0x17, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x32, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x48, 0x45, 0x49, 0x47, 0x48, 0x54, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10,
	0x09, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x5f, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x0c, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x55, 0x54, 0x54,
	0x4f, 0x4e, 0x10, 0x0d, 0x32, 0x97, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x0f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x1a,
	0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x54, 0x4c,
	0x69, 0x53, 0x75, 0x6e, 0x6e, 0x79, 0x69, 0x2f, 0x63, 0x69, 0x74, 0x61, 0x2d, 0x73, 0x64, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData = file_storage_proto_rawDesc
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_proto_rawDescData)
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storage_proto_goTypes = []interface{}{
	(Regions)(0),       // 0: storage.Regions
	(*Content)(nil),    // 1: storage.Content
	(*ExtKey)(nil),     // 2: storage.ExtKey
	(*Value)(nil),      // 3: storage.Value
	(*StatusCode)(nil), // 4: common.StatusCode
}
var file_storage_proto_depIdxs = []int32{
	4, // 0: storage.Value.status:type_name -> common.StatusCode
	1, // 1: storage.StorageService.Store:input_type -> storage.Content
	2, // 2: storage.StorageService.Load:input_type -> storage.ExtKey
	2, // 3: storage.StorageService.Delete:input_type -> storage.ExtKey
	4, // 4: storage.StorageService.Store:output_type -> common.StatusCode
	3, // 5: storage.StorageService.Load:output_type -> storage.Value
	4, // 6: storage.StorageService.Delete:output_type -> common.StatusCode
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtKey); i {
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
		file_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
			RawDescriptor: file_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		EnumInfos:         file_storage_proto_enumTypes,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_rawDesc = nil
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageServiceClient interface {
	// store key/value
	Store(ctx context.Context, in *Content, opts ...grpc.CallOption) (*StatusCode, error)
	// given a ext key return value
	Load(ctx context.Context, in *ExtKey, opts ...grpc.CallOption) (*Value, error)
	// given a ext key delete it
	Delete(ctx context.Context, in *ExtKey, opts ...grpc.CallOption) (*StatusCode, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) Store(ctx context.Context, in *Content, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := c.cc.Invoke(ctx, "/storage.StorageService/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) Load(ctx context.Context, in *ExtKey, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/storage.StorageService/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) Delete(ctx context.Context, in *ExtKey, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := c.cc.Invoke(ctx, "/storage.StorageService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
type StorageServiceServer interface {
	// store key/value
	Store(context.Context, *Content) (*StatusCode, error)
	// given a ext key return value
	Load(context.Context, *ExtKey) (*Value, error)
	// given a ext key delete it
	Delete(context.Context, *ExtKey) (*StatusCode, error)
}

// UnimplementedStorageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (*UnimplementedStorageServiceServer) Store(context.Context, *Content) (*StatusCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedStorageServiceServer) Load(context.Context, *ExtKey) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (*UnimplementedStorageServiceServer) Delete(context.Context, *ExtKey) (*StatusCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterStorageServiceServer(s *grpc.Server, srv StorageServiceServer) {
	s.RegisterService(&_StorageService_serviceDesc, srv)
}

func _StorageService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Store(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Load(ctx, req.(*ExtKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Delete(ctx, req.(*ExtKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _StorageService_Store_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _StorageService_Load_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StorageService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
