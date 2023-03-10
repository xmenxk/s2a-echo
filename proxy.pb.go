// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra_internal/appengine/playground-sheriffbot/appengine/helper/proxy.proto

package main

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Generic info request.
	Request *string `protobuf:"bytes,1,opt,name=request,proto3,oneof" json:"request,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *InfoRequest) GetRequest() string {
	if x != nil && x.Request != nil {
		return *x.Request
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The response information.
	Response *string `protobuf:"bytes,1,opt,name=response,proto3,oneof" json:"response,omitempty"`
	// Current version information.
	Version *int32 `protobuf:"varint,2,opt,name=version,proto3,oneof" json:"version,omitempty"`
	// Task URL or other server identifier if supported.
	Url *string `protobuf:"bytes,3,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetResponse() string {
	if x != nil && x.Response != nil {
		return *x.Response
	}
	return ""
}

func (x *InfoResponse) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *InfoResponse) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

var File_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto protoreflect.FileDescriptor

var file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x62, 0x6f, 0x74,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x86, 0x01,
	0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x32, 0x4f, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x46, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x41, 0x5a, 0x3f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x73,
	0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescOnce sync.Once
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescData = file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDesc
)

func file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescGZIP() []byte {
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescOnce.Do(func() {
		file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescData)
	})
	return file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDescData
}

var file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_goTypes = []interface{}{
	(*InfoRequest)(nil),  // 0: helper.InfoRequest
	(*InfoResponse)(nil), // 1: helper.InfoResponse
}
var file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_depIdxs = []int32{
	0, // 0: helper.Proxy.Info:input_type -> helper.InfoRequest
	1, // 1: helper.Proxy.Info:output_type -> helper.InfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_init() }
func file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_init() {
	if File_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_goTypes,
		DependencyIndexes: file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_depIdxs,
		MessageInfos:      file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_msgTypes,
	}.Build()
	File_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto = out.File
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_rawDesc = nil
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_goTypes = nil
	file_infra_internal_appengine_playground_sheriffbot_appengine_helper_proxy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyClient interface {
	// Returns information about service.
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}
type proxyPRPCClient struct {
	client *prpc.Client
}

func NewProxyPRPCClient(client *prpc.Client) ProxyClient {
	return &proxyPRPCClient{client}
}

func (c *proxyPRPCClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.client.Call(ctx, "cloud.alphabetcloud.directpath.Proxy", "Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/cloud.alphabetcloud.directpath.Proxy/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServer is the server API for Proxy service.
type ProxyServer interface {
	// Returns information about service.
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
}

// UnimplementedProxyServer can be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (*UnimplementedProxyServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterProxyServer(s prpc.Registrar, srv ProxyServer) {
	s.RegisterService(&_Proxy_serviceDesc, srv)
}

func _Proxy_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.alphabetcloud.directpath.Proxy/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Proxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.alphabetcloud.directpath.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Proxy_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra_internal/appengine/playground-sheriffbot/appengine/helper/proxy.proto",
}
