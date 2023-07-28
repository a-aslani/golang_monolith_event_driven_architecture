// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gamerspb/api.proto

package gamerspb

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
	GamersService_CreateGamer_FullMethodName     = "/gamerspb.GamersService/CreateGamer"
	GamersService_DisapproveGamer_FullMethodName = "/gamerspb.GamersService/DisapproveGamer"
	GamersService_GetGamer_FullMethodName        = "/gamerspb.GamersService/GetGamer"
	GamersService_GetGamers_FullMethodName       = "/gamerspb.GamersService/GetGamers"
)

// GamersServiceClient is the client API for GamersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GamersServiceClient interface {
	CreateGamer(ctx context.Context, in *CreateGamerRequest, opts ...grpc.CallOption) (*CreateGamerResponse, error)
	DisapproveGamer(ctx context.Context, in *DisapproveGamerRequest, opts ...grpc.CallOption) (*DisapproveGamerResponse, error)
	GetGamer(ctx context.Context, in *GetGamerRequest, opts ...grpc.CallOption) (*GetGamerResponse, error)
	GetGamers(ctx context.Context, in *GetGamersRequest, opts ...grpc.CallOption) (*GetGamersResponse, error)
}

type gamersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGamersServiceClient(cc grpc.ClientConnInterface) GamersServiceClient {
	return &gamersServiceClient{cc}
}

func (c *gamersServiceClient) CreateGamer(ctx context.Context, in *CreateGamerRequest, opts ...grpc.CallOption) (*CreateGamerResponse, error) {
	out := new(CreateGamerResponse)
	err := c.cc.Invoke(ctx, GamersService_CreateGamer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gamersServiceClient) DisapproveGamer(ctx context.Context, in *DisapproveGamerRequest, opts ...grpc.CallOption) (*DisapproveGamerResponse, error) {
	out := new(DisapproveGamerResponse)
	err := c.cc.Invoke(ctx, GamersService_DisapproveGamer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gamersServiceClient) GetGamer(ctx context.Context, in *GetGamerRequest, opts ...grpc.CallOption) (*GetGamerResponse, error) {
	out := new(GetGamerResponse)
	err := c.cc.Invoke(ctx, GamersService_GetGamer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gamersServiceClient) GetGamers(ctx context.Context, in *GetGamersRequest, opts ...grpc.CallOption) (*GetGamersResponse, error) {
	out := new(GetGamersResponse)
	err := c.cc.Invoke(ctx, GamersService_GetGamers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GamersServiceServer is the server API for GamersService service.
// All implementations must embed UnimplementedGamersServiceServer
// for forward compatibility
type GamersServiceServer interface {
	CreateGamer(context.Context, *CreateGamerRequest) (*CreateGamerResponse, error)
	DisapproveGamer(context.Context, *DisapproveGamerRequest) (*DisapproveGamerResponse, error)
	GetGamer(context.Context, *GetGamerRequest) (*GetGamerResponse, error)
	GetGamers(context.Context, *GetGamersRequest) (*GetGamersResponse, error)
	mustEmbedUnimplementedGamersServiceServer()
}

// UnimplementedGamersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGamersServiceServer struct {
}

func (UnimplementedGamersServiceServer) CreateGamer(context.Context, *CreateGamerRequest) (*CreateGamerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGamer not implemented")
}
func (UnimplementedGamersServiceServer) DisapproveGamer(context.Context, *DisapproveGamerRequest) (*DisapproveGamerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisapproveGamer not implemented")
}
func (UnimplementedGamersServiceServer) GetGamer(context.Context, *GetGamerRequest) (*GetGamerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGamer not implemented")
}
func (UnimplementedGamersServiceServer) GetGamers(context.Context, *GetGamersRequest) (*GetGamersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGamers not implemented")
}
func (UnimplementedGamersServiceServer) mustEmbedUnimplementedGamersServiceServer() {}

// UnsafeGamersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GamersServiceServer will
// result in compilation errors.
type UnsafeGamersServiceServer interface {
	mustEmbedUnimplementedGamersServiceServer()
}

func RegisterGamersServiceServer(s grpc.ServiceRegistrar, srv GamersServiceServer) {
	s.RegisterService(&GamersService_ServiceDesc, srv)
}

func _GamersService_CreateGamer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGamerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamersServiceServer).CreateGamer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GamersService_CreateGamer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamersServiceServer).CreateGamer(ctx, req.(*CreateGamerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GamersService_DisapproveGamer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisapproveGamerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamersServiceServer).DisapproveGamer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GamersService_DisapproveGamer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamersServiceServer).DisapproveGamer(ctx, req.(*DisapproveGamerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GamersService_GetGamer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGamerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamersServiceServer).GetGamer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GamersService_GetGamer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamersServiceServer).GetGamer(ctx, req.(*GetGamerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GamersService_GetGamers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGamersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamersServiceServer).GetGamers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GamersService_GetGamers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamersServiceServer).GetGamers(ctx, req.(*GetGamersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GamersService_ServiceDesc is the grpc.ServiceDesc for GamersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GamersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gamerspb.GamersService",
	HandlerType: (*GamersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGamer",
			Handler:    _GamersService_CreateGamer_Handler,
		},
		{
			MethodName: "DisapproveGamer",
			Handler:    _GamersService_DisapproveGamer_Handler,
		},
		{
			MethodName: "GetGamer",
			Handler:    _GamersService_GetGamer_Handler,
		},
		{
			MethodName: "GetGamers",
			Handler:    _GamersService_GetGamers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gamerspb/api.proto",
}
