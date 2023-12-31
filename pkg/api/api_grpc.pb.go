// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	PlayerCreate(ctx context.Context, in *PlayerCreateRequest, opts ...grpc.CallOption) (*PlayerCreateResponse, error)
	PlayerAsyncCreate(ctx context.Context, in *PlayerCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayerList(ctx context.Context, in *PlayerListRequest, opts ...grpc.CallOption) (*PlayerListResponse, error)
	PlayerPubsubList(ctx context.Context, in *PlayerListRequest, opts ...grpc.CallOption) (*PlayerListResponse, error)
	PlayerStreamList(ctx context.Context, in *PlayerStreamListRequest, opts ...grpc.CallOption) (Admin_PlayerStreamListClient, error)
	PlayerGet(ctx context.Context, in *PlayerGetRequest, opts ...grpc.CallOption) (*PlayerGetResponse, error)
	PlayerPubsubGet(ctx context.Context, in *PlayerGetRequest, opts ...grpc.CallOption) (*PlayerGetResponse, error)
	PlayerUpdate(ctx context.Context, in *PlayerUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayerAsyncUpdate(ctx context.Context, in *PlayerUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayerDelete(ctx context.Context, in *PlayerDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayerAsyncDelete(ctx context.Context, in *PlayerDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) PlayerCreate(ctx context.Context, in *PlayerCreateRequest, opts ...grpc.CallOption) (*PlayerCreateResponse, error) {
	out := new(PlayerCreateResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerAsyncCreate(ctx context.Context, in *PlayerCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerAsyncCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerList(ctx context.Context, in *PlayerListRequest, opts ...grpc.CallOption) (*PlayerListResponse, error) {
	out := new(PlayerListResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerPubsubList(ctx context.Context, in *PlayerListRequest, opts ...grpc.CallOption) (*PlayerListResponse, error) {
	out := new(PlayerListResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerPubsubList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerStreamList(ctx context.Context, in *PlayerStreamListRequest, opts ...grpc.CallOption) (Admin_PlayerStreamListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Admin_ServiceDesc.Streams[0], "/ozon.dev.mc2.api.Admin/PlayerStreamList", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminPlayerStreamListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_PlayerStreamListClient interface {
	Recv() (*PlayerStreamListResponse, error)
	grpc.ClientStream
}

type adminPlayerStreamListClient struct {
	grpc.ClientStream
}

func (x *adminPlayerStreamListClient) Recv() (*PlayerStreamListResponse, error) {
	m := new(PlayerStreamListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminClient) PlayerGet(ctx context.Context, in *PlayerGetRequest, opts ...grpc.CallOption) (*PlayerGetResponse, error) {
	out := new(PlayerGetResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerPubsubGet(ctx context.Context, in *PlayerGetRequest, opts ...grpc.CallOption) (*PlayerGetResponse, error) {
	out := new(PlayerGetResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerPubsubGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerUpdate(ctx context.Context, in *PlayerUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerAsyncUpdate(ctx context.Context, in *PlayerUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerAsyncUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerDelete(ctx context.Context, in *PlayerDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) PlayerAsyncDelete(ctx context.Context, in *PlayerDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozon.dev.mc2.api.Admin/PlayerAsyncDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	PlayerCreate(context.Context, *PlayerCreateRequest) (*PlayerCreateResponse, error)
	PlayerAsyncCreate(context.Context, *PlayerCreateRequest) (*emptypb.Empty, error)
	PlayerList(context.Context, *PlayerListRequest) (*PlayerListResponse, error)
	PlayerPubsubList(context.Context, *PlayerListRequest) (*PlayerListResponse, error)
	PlayerStreamList(*PlayerStreamListRequest, Admin_PlayerStreamListServer) error
	PlayerGet(context.Context, *PlayerGetRequest) (*PlayerGetResponse, error)
	PlayerPubsubGet(context.Context, *PlayerGetRequest) (*PlayerGetResponse, error)
	PlayerUpdate(context.Context, *PlayerUpdateRequest) (*emptypb.Empty, error)
	PlayerAsyncUpdate(context.Context, *PlayerUpdateRequest) (*emptypb.Empty, error)
	PlayerDelete(context.Context, *PlayerDeleteRequest) (*emptypb.Empty, error)
	PlayerAsyncDelete(context.Context, *PlayerDeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) PlayerCreate(context.Context, *PlayerCreateRequest) (*PlayerCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerCreate not implemented")
}
func (UnimplementedAdminServer) PlayerAsyncCreate(context.Context, *PlayerCreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerAsyncCreate not implemented")
}
func (UnimplementedAdminServer) PlayerList(context.Context, *PlayerListRequest) (*PlayerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerList not implemented")
}
func (UnimplementedAdminServer) PlayerPubsubList(context.Context, *PlayerListRequest) (*PlayerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerPubsubList not implemented")
}
func (UnimplementedAdminServer) PlayerStreamList(*PlayerStreamListRequest, Admin_PlayerStreamListServer) error {
	return status.Errorf(codes.Unimplemented, "method PlayerStreamList not implemented")
}
func (UnimplementedAdminServer) PlayerGet(context.Context, *PlayerGetRequest) (*PlayerGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerGet not implemented")
}
func (UnimplementedAdminServer) PlayerPubsubGet(context.Context, *PlayerGetRequest) (*PlayerGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerPubsubGet not implemented")
}
func (UnimplementedAdminServer) PlayerUpdate(context.Context, *PlayerUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerUpdate not implemented")
}
func (UnimplementedAdminServer) PlayerAsyncUpdate(context.Context, *PlayerUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerAsyncUpdate not implemented")
}
func (UnimplementedAdminServer) PlayerDelete(context.Context, *PlayerDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerDelete not implemented")
}
func (UnimplementedAdminServer) PlayerAsyncDelete(context.Context, *PlayerDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerAsyncDelete not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_PlayerCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerCreate(ctx, req.(*PlayerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerAsyncCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerAsyncCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerAsyncCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerAsyncCreate(ctx, req.(*PlayerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerList(ctx, req.(*PlayerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerPubsubList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerPubsubList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerPubsubList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerPubsubList(ctx, req.(*PlayerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerStreamList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlayerStreamListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).PlayerStreamList(m, &adminPlayerStreamListServer{stream})
}

type Admin_PlayerStreamListServer interface {
	Send(*PlayerStreamListResponse) error
	grpc.ServerStream
}

type adminPlayerStreamListServer struct {
	grpc.ServerStream
}

func (x *adminPlayerStreamListServer) Send(m *PlayerStreamListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Admin_PlayerGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerGet(ctx, req.(*PlayerGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerPubsubGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerPubsubGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerPubsubGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerPubsubGet(ctx, req.(*PlayerGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerUpdate(ctx, req.(*PlayerUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerAsyncUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerAsyncUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerAsyncUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerAsyncUpdate(ctx, req.(*PlayerUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerDelete(ctx, req.(*PlayerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_PlayerAsyncDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).PlayerAsyncDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.mc2.api.Admin/PlayerAsyncDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).PlayerAsyncDelete(ctx, req.(*PlayerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozon.dev.mc2.api.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayerCreate",
			Handler:    _Admin_PlayerCreate_Handler,
		},
		{
			MethodName: "PlayerAsyncCreate",
			Handler:    _Admin_PlayerAsyncCreate_Handler,
		},
		{
			MethodName: "PlayerList",
			Handler:    _Admin_PlayerList_Handler,
		},
		{
			MethodName: "PlayerPubsubList",
			Handler:    _Admin_PlayerPubsubList_Handler,
		},
		{
			MethodName: "PlayerGet",
			Handler:    _Admin_PlayerGet_Handler,
		},
		{
			MethodName: "PlayerPubsubGet",
			Handler:    _Admin_PlayerPubsubGet_Handler,
		},
		{
			MethodName: "PlayerUpdate",
			Handler:    _Admin_PlayerUpdate_Handler,
		},
		{
			MethodName: "PlayerAsyncUpdate",
			Handler:    _Admin_PlayerAsyncUpdate_Handler,
		},
		{
			MethodName: "PlayerDelete",
			Handler:    _Admin_PlayerDelete_Handler,
		},
		{
			MethodName: "PlayerAsyncDelete",
			Handler:    _Admin_PlayerAsyncDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PlayerStreamList",
			Handler:       _Admin_PlayerStreamList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
