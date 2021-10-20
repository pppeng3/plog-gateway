// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plog_gateway

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

// PLogGatewayClient is the client API for PLogGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PLogGatewayClient interface {
	UploadLog(ctx context.Context, in *UploadLogRequest, opts ...grpc.CallOption) (*UploadLogResponse, error)
}

type pLogGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewPLogGatewayClient(cc grpc.ClientConnInterface) PLogGatewayClient {
	return &pLogGatewayClient{cc}
}

func (c *pLogGatewayClient) UploadLog(ctx context.Context, in *UploadLogRequest, opts ...grpc.CallOption) (*UploadLogResponse, error) {
	out := new(UploadLogResponse)
	err := c.cc.Invoke(ctx, "/plog_gateway.PLogGateway/UploadLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PLogGatewayServer is the server API for PLogGateway service.
// All implementations must embed UnimplementedPLogGatewayServer
// for forward compatibility
type PLogGatewayServer interface {
	UploadLog(context.Context, *UploadLogRequest) (*UploadLogResponse, error)
}

// UnimplementedPLogGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedPLogGatewayServer struct {
}

func (UnimplementedPLogGatewayServer) UploadLog(context.Context, *UploadLogRequest) (*UploadLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadLog not implemented")
}
func (UnimplementedPLogGatewayServer) mustEmbedUnimplementedPLogGatewayServer() {}

// UnsafePLogGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PLogGatewayServer will
// result in compilation errors.
type UnsafePLogGatewayServer interface {
	mustEmbedUnimplementedPLogGatewayServer()
}

func RegisterPLogGatewayServer(s grpc.ServiceRegistrar, srv PLogGatewayServer) {
	s.RegisterService(&PLogGateway_ServiceDesc, srv)
}

func _PLogGateway_UploadLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PLogGatewayServer).UploadLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plog_gateway.PLogGateway/UploadLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PLogGatewayServer).UploadLog(ctx, req.(*UploadLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PLogGateway_ServiceDesc is the grpc.ServiceDesc for PLogGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PLogGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plog_gateway.PLogGateway",
	HandlerType: (*PLogGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadLog",
			Handler:    _PLogGateway_UploadLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/plog_gateway.proto",
}
