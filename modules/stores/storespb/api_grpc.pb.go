// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: storespb/api.proto

package storespb

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

const (
	StoresService_Test_FullMethodName = "/storespb.StoresService/Test"
)

// StoresServiceClient is the client API for StoresService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoresServiceClient interface {
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type storesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoresServiceClient(cc grpc.ClientConnInterface) StoresServiceClient {
	return &storesServiceClient{cc}
}

func (c *storesServiceClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, StoresService_Test_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoresServiceServer is the server API for StoresService service.
// All implementations must embed UnimplementedStoresServiceServer
// for forward compatibility
type StoresServiceServer interface {
	Test(context.Context, *TestRequest) (*TestResponse, error)
	mustEmbedUnimplementedStoresServiceServer()
}

// UnimplementedStoresServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoresServiceServer struct {
}

func (UnimplementedStoresServiceServer) Test(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedStoresServiceServer) mustEmbedUnimplementedStoresServiceServer() {}

// UnsafeStoresServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoresServiceServer will
// result in compilation errors.
type UnsafeStoresServiceServer interface {
	mustEmbedUnimplementedStoresServiceServer()
}

func RegisterStoresServiceServer(s grpc.ServiceRegistrar, srv StoresServiceServer) {
	s.RegisterService(&StoresService_ServiceDesc, srv)
}

func _StoresService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoresService_Test_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoresService_ServiceDesc is the grpc.ServiceDesc for StoresService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoresService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storespb.StoresService",
	HandlerType: (*StoresServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _StoresService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storespb/api.proto",
}