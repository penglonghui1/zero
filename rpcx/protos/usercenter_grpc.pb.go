// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: usercenter.proto

package pb

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

// OkrClient is the client API for Okr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OkrClient interface {
	GetOkr(ctx context.Context, in *GetOkrRequest, opts ...grpc.CallOption) (*Response, error)
}

type okrClient struct {
	cc grpc.ClientConnInterface
}

func NewOkrClient(cc grpc.ClientConnInterface) OkrClient {
	return &okrClient{cc}
}

func (c *okrClient) GetOkr(ctx context.Context, in *GetOkrRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.Okr/GetOkr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OkrServer is the server API for Okr service.
// All implementations must embed UnimplementedOkrServer
// for forward compatibility
type OkrServer interface {
	GetOkr(context.Context, *GetOkrRequest) (*Response, error)
	mustEmbedUnimplementedOkrServer()
}

// UnimplementedOkrServer must be embedded to have forward compatible implementations.
type UnimplementedOkrServer struct {
}

func (UnimplementedOkrServer) GetOkr(context.Context, *GetOkrRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOkr not implemented")
}
func (UnimplementedOkrServer) mustEmbedUnimplementedOkrServer() {}

// UnsafeOkrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OkrServer will
// result in compilation errors.
type UnsafeOkrServer interface {
	mustEmbedUnimplementedOkrServer()
}

func RegisterOkrServer(s grpc.ServiceRegistrar, srv OkrServer) {
	s.RegisterService(&Okr_ServiceDesc, srv)
}

func _Okr_GetOkr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOkrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrServer).GetOkr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Okr/GetOkr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrServer).GetOkr(ctx, req.(*GetOkrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Okr_ServiceDesc is the grpc.ServiceDesc for Okr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Okr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Okr",
	HandlerType: (*OkrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOkr",
			Handler:    _Okr_GetOkr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}
