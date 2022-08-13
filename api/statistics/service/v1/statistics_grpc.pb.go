// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/statistics/service/v1/statistics.proto

package v1

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

// StatisticsServiceClient is the client API for StatisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticsServiceClient interface {
	GetUserStatistics(ctx context.Context, in *GetUserStatisticsRequest, opts ...grpc.CallOption) (*GetUserStatisticsReply, error)
	CreateStatistics(ctx context.Context, in *CreateStatisticsRequest, opts ...grpc.CallOption) (*CreateStatisticsReply, error)
}

type statisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticsServiceClient(cc grpc.ClientConnInterface) StatisticsServiceClient {
	return &statisticsServiceClient{cc}
}

func (c *statisticsServiceClient) GetUserStatistics(ctx context.Context, in *GetUserStatisticsRequest, opts ...grpc.CallOption) (*GetUserStatisticsReply, error) {
	out := new(GetUserStatisticsReply)
	err := c.cc.Invoke(ctx, "/api.statistics.service.v1.StatisticsService/GetUserStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsServiceClient) CreateStatistics(ctx context.Context, in *CreateStatisticsRequest, opts ...grpc.CallOption) (*CreateStatisticsReply, error) {
	out := new(CreateStatisticsReply)
	err := c.cc.Invoke(ctx, "/api.statistics.service.v1.StatisticsService/CreateStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsServiceServer is the server API for StatisticsService service.
// All implementations must embed UnimplementedStatisticsServiceServer
// for forward compatibility
type StatisticsServiceServer interface {
	GetUserStatistics(context.Context, *GetUserStatisticsRequest) (*GetUserStatisticsReply, error)
	CreateStatistics(context.Context, *CreateStatisticsRequest) (*CreateStatisticsReply, error)
	mustEmbedUnimplementedStatisticsServiceServer()
}

// UnimplementedStatisticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticsServiceServer struct {
}

func (UnimplementedStatisticsServiceServer) GetUserStatistics(context.Context, *GetUserStatisticsRequest) (*GetUserStatisticsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStatistics not implemented")
}
func (UnimplementedStatisticsServiceServer) CreateStatistics(context.Context, *CreateStatisticsRequest) (*CreateStatisticsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatistics not implemented")
}
func (UnimplementedStatisticsServiceServer) mustEmbedUnimplementedStatisticsServiceServer() {}

// UnsafeStatisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticsServiceServer will
// result in compilation errors.
type UnsafeStatisticsServiceServer interface {
	mustEmbedUnimplementedStatisticsServiceServer()
}

func RegisterStatisticsServiceServer(s grpc.ServiceRegistrar, srv StatisticsServiceServer) {
	s.RegisterService(&StatisticsService_ServiceDesc, srv)
}

func _StatisticsService_GetUserStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServiceServer).GetUserStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.statistics.service.v1.StatisticsService/GetUserStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServiceServer).GetUserStatistics(ctx, req.(*GetUserStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatisticsService_CreateStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServiceServer).CreateStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.statistics.service.v1.StatisticsService/CreateStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServiceServer).CreateStatistics(ctx, req.(*CreateStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatisticsService_ServiceDesc is the grpc.ServiceDesc for StatisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.statistics.service.v1.StatisticsService",
	HandlerType: (*StatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserStatistics",
			Handler:    _StatisticsService_GetUserStatistics_Handler,
		},
		{
			MethodName: "CreateStatistics",
			Handler:    _StatisticsService_CreateStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/statistics/service/v1/statistics.proto",
}
