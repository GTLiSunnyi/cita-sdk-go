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
// source: protos/consensus.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_protos_consensus_proto protoreflect.FileDescriptor

var file_protos_consensus_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x92, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x0b, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x10, 0x5a,
	0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_consensus_proto_goTypes = []interface{}{
	(*ConsensusConfiguration)(nil), // 0: common.ConsensusConfiguration
	(*ProposalWithProof)(nil),      // 1: common.ProposalWithProof
	(*StatusCode)(nil),             // 2: common.StatusCode
}
var file_protos_consensus_proto_depIdxs = []int32{
	0, // 0: consensus.ConsensusService.Reconfigure:input_type -> common.ConsensusConfiguration
	1, // 1: consensus.ConsensusService.CheckBlock:input_type -> common.ProposalWithProof
	2, // 2: consensus.ConsensusService.Reconfigure:output_type -> common.StatusCode
	2, // 3: consensus.ConsensusService.CheckBlock:output_type -> common.StatusCode
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_consensus_proto_init() }
func file_protos_consensus_proto_init() {
	if File_protos_consensus_proto != nil {
		return
	}
	file_protos_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_consensus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_consensus_proto_goTypes,
		DependencyIndexes: file_protos_consensus_proto_depIdxs,
	}.Build()
	File_protos_consensus_proto = out.File
	file_protos_consensus_proto_rawDesc = nil
	file_protos_consensus_proto_goTypes = nil
	file_protos_consensus_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsensusServiceClient is the client API for ConsensusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsensusServiceClient interface {
	/// reconfigure consensus status
	Reconfigure(ctx context.Context, in *ConsensusConfiguration, opts ...grpc.CallOption) (*StatusCode, error)
	/// check block validity
	CheckBlock(ctx context.Context, in *ProposalWithProof, opts ...grpc.CallOption) (*StatusCode, error)
}

type consensusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsensusServiceClient(cc grpc.ClientConnInterface) ConsensusServiceClient {
	return &consensusServiceClient{cc}
}

func (c *consensusServiceClient) Reconfigure(ctx context.Context, in *ConsensusConfiguration, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := c.cc.Invoke(ctx, "/consensus.ConsensusService/Reconfigure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) CheckBlock(ctx context.Context, in *ProposalWithProof, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := c.cc.Invoke(ctx, "/consensus.ConsensusService/CheckBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsensusServiceServer is the server API for ConsensusService service.
type ConsensusServiceServer interface {
	/// reconfigure consensus status
	Reconfigure(context.Context, *ConsensusConfiguration) (*StatusCode, error)
	/// check block validity
	CheckBlock(context.Context, *ProposalWithProof) (*StatusCode, error)
}

// UnimplementedConsensusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsensusServiceServer struct {
}

func (*UnimplementedConsensusServiceServer) Reconfigure(context.Context, *ConsensusConfiguration) (*StatusCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reconfigure not implemented")
}
func (*UnimplementedConsensusServiceServer) CheckBlock(context.Context, *ProposalWithProof) (*StatusCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBlock not implemented")
}

func RegisterConsensusServiceServer(s *grpc.Server, srv ConsensusServiceServer) {
	s.RegisterService(&_ConsensusService_serviceDesc, srv)
}

func _ConsensusService_Reconfigure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsensusConfiguration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).Reconfigure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consensus.ConsensusService/Reconfigure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).Reconfigure(ctx, req.(*ConsensusConfiguration))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_CheckBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalWithProof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).CheckBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consensus.ConsensusService/CheckBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).CheckBlock(ctx, req.(*ProposalWithProof))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsensusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consensus.ConsensusService",
	HandlerType: (*ConsensusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reconfigure",
			Handler:    _ConsensusService_Reconfigure_Handler,
		},
		{
			MethodName: "CheckBlock",
			Handler:    _ConsensusService_CheckBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/consensus.proto",
}
