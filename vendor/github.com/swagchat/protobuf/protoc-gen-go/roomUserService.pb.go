// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roomUserService.proto

package protoc_gen_go

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RoomUserService service

type RoomUserServiceClient interface {
	AddRoomUsers(ctx context.Context, in *AddRoomUsersRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	RetrieveRoomUsers(ctx context.Context, in *RetrieveRoomUsersRequest, opts ...grpc.CallOption) (*RoomUsersResponse, error)
	RetrieveRoomUserIds(ctx context.Context, in *RetrieveRoomUsersRequest, opts ...grpc.CallOption) (*RoomUserIdsResponse, error)
	UpdateRoomUser(ctx context.Context, in *UpdateRoomUserRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	DeleteRoomUsers(ctx context.Context, in *DeleteRoomUsersRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type roomUserServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoomUserServiceClient(cc *grpc.ClientConn) RoomUserServiceClient {
	return &roomUserServiceClient{cc}
}

func (c *roomUserServiceClient) AddRoomUsers(ctx context.Context, in *AddRoomUsersRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.RoomUserService/AddRoomUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomUserServiceClient) RetrieveRoomUsers(ctx context.Context, in *RetrieveRoomUsersRequest, opts ...grpc.CallOption) (*RoomUsersResponse, error) {
	out := new(RoomUsersResponse)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.RoomUserService/RetrieveRoomUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomUserServiceClient) RetrieveRoomUserIds(ctx context.Context, in *RetrieveRoomUsersRequest, opts ...grpc.CallOption) (*RoomUserIdsResponse, error) {
	out := new(RoomUserIdsResponse)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.RoomUserService/RetrieveRoomUserIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomUserServiceClient) UpdateRoomUser(ctx context.Context, in *UpdateRoomUserRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.RoomUserService/UpdateRoomUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomUserServiceClient) DeleteRoomUsers(ctx context.Context, in *DeleteRoomUsersRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.RoomUserService/DeleteRoomUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomUserService service

type RoomUserServiceServer interface {
	AddRoomUsers(context.Context, *AddRoomUsersRequest) (*google_protobuf1.Empty, error)
	RetrieveRoomUsers(context.Context, *RetrieveRoomUsersRequest) (*RoomUsersResponse, error)
	RetrieveRoomUserIds(context.Context, *RetrieveRoomUsersRequest) (*RoomUserIdsResponse, error)
	UpdateRoomUser(context.Context, *UpdateRoomUserRequest) (*google_protobuf1.Empty, error)
	DeleteRoomUsers(context.Context, *DeleteRoomUsersRequest) (*google_protobuf1.Empty, error)
}

func RegisterRoomUserServiceServer(s *grpc.Server, srv RoomUserServiceServer) {
	s.RegisterService(&_RoomUserService_serviceDesc, srv)
}

func _RoomUserService_AddRoomUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoomUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomUserServiceServer).AddRoomUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.RoomUserService/AddRoomUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomUserServiceServer).AddRoomUsers(ctx, req.(*AddRoomUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomUserService_RetrieveRoomUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRoomUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomUserServiceServer).RetrieveRoomUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.RoomUserService/RetrieveRoomUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomUserServiceServer).RetrieveRoomUsers(ctx, req.(*RetrieveRoomUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomUserService_RetrieveRoomUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRoomUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomUserServiceServer).RetrieveRoomUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.RoomUserService/RetrieveRoomUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomUserServiceServer).RetrieveRoomUserIds(ctx, req.(*RetrieveRoomUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomUserService_UpdateRoomUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomUserServiceServer).UpdateRoomUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.RoomUserService/UpdateRoomUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomUserServiceServer).UpdateRoomUser(ctx, req.(*UpdateRoomUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomUserService_DeleteRoomUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomUserServiceServer).DeleteRoomUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.RoomUserService/DeleteRoomUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomUserServiceServer).DeleteRoomUsers(ctx, req.(*DeleteRoomUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomUserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "swagchat.protobuf.RoomUserService",
	HandlerType: (*RoomUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRoomUsers",
			Handler:    _RoomUserService_AddRoomUsers_Handler,
		},
		{
			MethodName: "RetrieveRoomUsers",
			Handler:    _RoomUserService_RetrieveRoomUsers_Handler,
		},
		{
			MethodName: "RetrieveRoomUserIds",
			Handler:    _RoomUserService_RetrieveRoomUserIds_Handler,
		},
		{
			MethodName: "UpdateRoomUser",
			Handler:    _RoomUserService_UpdateRoomUser_Handler,
		},
		{
			MethodName: "DeleteRoomUsers",
			Handler:    _RoomUserService_DeleteRoomUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roomUserService.proto",
}

func init() { proto.RegisterFile("roomUserService.proto", fileDescriptorRoomUserService) }

var fileDescriptorRoomUserService = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x0b, 0x82, 0x87, 0x45, 0x2c, 0x5d, 0xd1, 0x43, 0xf5, 0x54, 0xa4, 0xf8, 0xaf, 0x09,
	0xe8, 0x13, 0x28, 0x7a, 0xe8, 0x41, 0x90, 0x68, 0x3d, 0x78, 0xdb, 0x64, 0xc7, 0xcd, 0x42, 0xb2,
	0xb3, 0xee, 0x4e, 0x2a, 0x3e, 0xb2, 0x6f, 0x21, 0x36, 0x49, 0x93, 0xb6, 0xc1, 0xe2, 0x6d, 0x67,
	0xe7, 0xdb, 0xef, 0x37, 0x0c, 0xcb, 0x0e, 0x1d, 0x62, 0x3e, 0xf3, 0xe0, 0x9e, 0xc1, 0xcd, 0x75,
	0x02, 0x81, 0x75, 0x48, 0xc8, 0x07, 0xfe, 0x53, 0xa8, 0x24, 0x15, 0x54, 0xd6, 0x71, 0xf1, 0x3e,
	0x3c, 0x51, 0x88, 0x2a, 0x83, 0x50, 0x58, 0x1d, 0x0a, 0x63, 0x90, 0x04, 0x69, 0x34, 0xbe, 0x04,
	0x86, 0xc7, 0x55, 0xb7, 0xc6, 0x43, 0xc8, 0x2d, 0x7d, 0x55, 0xcd, 0x65, 0xc8, 0x23, 0x78, 0x2f,
	0x54, 0x15, 0x72, 0xfd, 0xbd, 0xc3, 0xfa, 0xd1, 0x6a, 0x3c, 0x7f, 0x62, 0x7b, 0xb7, 0x52, 0xd6,
	0xb7, 0x9e, 0x8f, 0x83, 0x8d, 0x49, 0x82, 0x36, 0x10, 0xc1, 0x47, 0x01, 0x9e, 0x86, 0x47, 0x41,
	0x39, 0x40, 0x43, 0x3d, 0xfc, 0x0e, 0x30, 0xea, 0xf1, 0x94, 0x0d, 0x22, 0x20, 0xa7, 0x61, 0x0e,
	0x8d, 0xf6, 0xb2, 0x43, 0xbb, 0x41, 0xd5, 0xee, 0xd3, 0x2e, 0xb8, 0x81, 0xbc, 0x45, 0xe3, 0x61,
	0xd4, 0xe3, 0x19, 0x3b, 0x58, 0x77, 0x4c, 0xe5, 0x3f, 0xb3, 0xc6, 0x7f, 0x64, 0x4d, 0x65, 0x3b,
	0xed, 0x85, 0xed, 0xcf, 0xac, 0x14, 0xb4, 0x74, 0xf0, 0xb3, 0x8e, 0xb7, 0xab, 0xc8, 0xf6, 0x6d,
	0xbd, 0xb2, 0xfe, 0x3d, 0x64, 0x40, 0xad, 0x5d, 0x9d, 0x77, 0x68, 0xd7, 0x98, 0xad, 0xde, 0xbb,
	0xab, 0xb7, 0x0b, 0xa5, 0x29, 0x2d, 0xe2, 0x20, 0xc1, 0x3c, 0xac, 0x85, 0xcd, 0x77, 0x59, 0x1c,
	0x92, 0x89, 0x02, 0x33, 0x51, 0x18, 0xef, 0x2e, 0xca, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa2, 0xb4, 0x2f, 0x36, 0x9e, 0x02, 0x00, 0x00,
}
