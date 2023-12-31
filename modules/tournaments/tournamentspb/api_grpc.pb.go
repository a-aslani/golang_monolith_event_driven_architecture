// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tournamentspb/api.proto

package tournamentspb

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
	TournamentsService_CreateTournament_FullMethodName = "/tournamentspb.TournamentsService/CreateTournament"
)

// TournamentsServiceClient is the client API for TournamentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TournamentsServiceClient interface {
	CreateTournament(ctx context.Context, in *CreateTournamentRequest, opts ...grpc.CallOption) (*CreateTournamentResponse, error)
}

type tournamentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTournamentsServiceClient(cc grpc.ClientConnInterface) TournamentsServiceClient {
	return &tournamentsServiceClient{cc}
}

func (c *tournamentsServiceClient) CreateTournament(ctx context.Context, in *CreateTournamentRequest, opts ...grpc.CallOption) (*CreateTournamentResponse, error) {
	out := new(CreateTournamentResponse)
	err := c.cc.Invoke(ctx, TournamentsService_CreateTournament_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TournamentsServiceServer is the server API for TournamentsService service.
// All implementations must embed UnimplementedTournamentsServiceServer
// for forward compatibility
type TournamentsServiceServer interface {
	CreateTournament(context.Context, *CreateTournamentRequest) (*CreateTournamentResponse, error)
	mustEmbedUnimplementedTournamentsServiceServer()
}

// UnimplementedTournamentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTournamentsServiceServer struct {
}

func (UnimplementedTournamentsServiceServer) CreateTournament(context.Context, *CreateTournamentRequest) (*CreateTournamentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTournament not implemented")
}
func (UnimplementedTournamentsServiceServer) mustEmbedUnimplementedTournamentsServiceServer() {}

// UnsafeTournamentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TournamentsServiceServer will
// result in compilation errors.
type UnsafeTournamentsServiceServer interface {
	mustEmbedUnimplementedTournamentsServiceServer()
}

func RegisterTournamentsServiceServer(s grpc.ServiceRegistrar, srv TournamentsServiceServer) {
	s.RegisterService(&TournamentsService_ServiceDesc, srv)
}

func _TournamentsService_CreateTournament_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTournamentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TournamentsServiceServer).CreateTournament(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TournamentsService_CreateTournament_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TournamentsServiceServer).CreateTournament(ctx, req.(*CreateTournamentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TournamentsService_ServiceDesc is the grpc.ServiceDesc for TournamentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TournamentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tournamentspb.TournamentsService",
	HandlerType: (*TournamentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTournament",
			Handler:    _TournamentsService_CreateTournament_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tournamentspb/api.proto",
}
