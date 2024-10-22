// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: yara.proto

package grpcyara

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// YARACheckClient is the client API for YARACheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YARACheckClient interface {
	YARACheck(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckReply, error)
	YARAFileCheck(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileReply, error)
}

type yARACheckClient struct {
	cc grpc.ClientConnInterface
}

func NewYARACheckClient(cc grpc.ClientConnInterface) YARACheckClient {
	return &yARACheckClient{cc}
}

func (c *yARACheckClient) YARACheck(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckReply, error) {
	out := new(CheckReply)
	err := c.cc.Invoke(ctx, "/YARACheck/YARACheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yARACheckClient) YARAFileCheck(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileReply, error) {
	out := new(FileReply)
	err := c.cc.Invoke(ctx, "/YARACheck/YARAFileCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YARACheckServer is the server API for YARACheck service.
// All implementations must embed UnimplementedYARACheckServer
// for forward compatibility
type YARACheckServer interface {
	YARACheck(context.Context, *CheckRequest) (*CheckReply, error)
	YARAFileCheck(context.Context, *FileRequest) (*FileReply, error)
	mustEmbedUnimplementedYARACheckServer()
}

// UnimplementedYARACheckServer must be embedded to have forward compatible implementations.
type UnimplementedYARACheckServer struct {
}

func (UnimplementedYARACheckServer) YARACheck(context.Context, *CheckRequest) (*CheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method YARACheck not implemented")
}
func (UnimplementedYARACheckServer) YARAFileCheck(context.Context, *FileRequest) (*FileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method YARAFileCheck not implemented")
}
func (UnimplementedYARACheckServer) mustEmbedUnimplementedYARACheckServer() {}

// UnsafeYARACheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YARACheckServer will
// result in compilation errors.
type UnsafeYARACheckServer interface {
	mustEmbedUnimplementedYARACheckServer()
}

func RegisterYARACheckServer(s grpc.ServiceRegistrar, srv YARACheckServer) {
	s.RegisterService(&YARACheck_ServiceDesc, srv)
}

func _YARACheck_YARACheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YARACheckServer).YARACheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YARACheck/YARACheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YARACheckServer).YARACheck(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YARACheck_YARAFileCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YARACheckServer).YARAFileCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YARACheck/YARAFileCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YARACheckServer).YARAFileCheck(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YARACheck_ServiceDesc is the grpc.ServiceDesc for YARACheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YARACheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YARACheck",
	HandlerType: (*YARACheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "YARACheck",
			Handler:    _YARACheck_YARACheck_Handler,
		},
		{
			MethodName: "YARAFileCheck",
			Handler:    _YARACheck_YARAFileCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yara.proto",
}
