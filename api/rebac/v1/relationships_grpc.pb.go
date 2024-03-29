// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rebac/v1/relationships.proto

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

const (
	Relationships_CreateRelationships_FullMethodName = "/api.rebac.v1.Relationships/CreateRelationships"
	Relationships_ReadRelationships_FullMethodName   = "/api.rebac.v1.Relationships/ReadRelationships"
	Relationships_DeleteRelationships_FullMethodName = "/api.rebac.v1.Relationships/DeleteRelationships"
)

// RelationshipsClient is the client API for Relationships service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationshipsClient interface {
	CreateRelationships(ctx context.Context, in *CreateRelationshipsRequest, opts ...grpc.CallOption) (*CreateRelationshipsResponse, error)
	ReadRelationships(ctx context.Context, in *ReadRelationshipsRequest, opts ...grpc.CallOption) (*ReadRelationshipsResponse, error)
	DeleteRelationships(ctx context.Context, in *DeleteRelationshipsRequest, opts ...grpc.CallOption) (*DeleteRelationshipsResponse, error)
}

type relationshipsClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationshipsClient(cc grpc.ClientConnInterface) RelationshipsClient {
	return &relationshipsClient{cc}
}

func (c *relationshipsClient) CreateRelationships(ctx context.Context, in *CreateRelationshipsRequest, opts ...grpc.CallOption) (*CreateRelationshipsResponse, error) {
	out := new(CreateRelationshipsResponse)
	err := c.cc.Invoke(ctx, Relationships_CreateRelationships_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipsClient) ReadRelationships(ctx context.Context, in *ReadRelationshipsRequest, opts ...grpc.CallOption) (*ReadRelationshipsResponse, error) {
	out := new(ReadRelationshipsResponse)
	err := c.cc.Invoke(ctx, Relationships_ReadRelationships_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipsClient) DeleteRelationships(ctx context.Context, in *DeleteRelationshipsRequest, opts ...grpc.CallOption) (*DeleteRelationshipsResponse, error) {
	out := new(DeleteRelationshipsResponse)
	err := c.cc.Invoke(ctx, Relationships_DeleteRelationships_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationshipsServer is the server API for Relationships service.
// All implementations must embed UnimplementedRelationshipsServer
// for forward compatibility
type RelationshipsServer interface {
	CreateRelationships(context.Context, *CreateRelationshipsRequest) (*CreateRelationshipsResponse, error)
	ReadRelationships(context.Context, *ReadRelationshipsRequest) (*ReadRelationshipsResponse, error)
	DeleteRelationships(context.Context, *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error)
	mustEmbedUnimplementedRelationshipsServer()
}

// UnimplementedRelationshipsServer must be embedded to have forward compatible implementations.
type UnimplementedRelationshipsServer struct {
}

func (UnimplementedRelationshipsServer) CreateRelationships(context.Context, *CreateRelationshipsRequest) (*CreateRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelationships not implemented")
}
func (UnimplementedRelationshipsServer) ReadRelationships(context.Context, *ReadRelationshipsRequest) (*ReadRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRelationships not implemented")
}
func (UnimplementedRelationshipsServer) DeleteRelationships(context.Context, *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationships not implemented")
}
func (UnimplementedRelationshipsServer) mustEmbedUnimplementedRelationshipsServer() {}

// UnsafeRelationshipsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationshipsServer will
// result in compilation errors.
type UnsafeRelationshipsServer interface {
	mustEmbedUnimplementedRelationshipsServer()
}

func RegisterRelationshipsServer(s grpc.ServiceRegistrar, srv RelationshipsServer) {
	s.RegisterService(&Relationships_ServiceDesc, srv)
}

func _Relationships_CreateRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipsServer).CreateRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relationships_CreateRelationships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipsServer).CreateRelationships(ctx, req.(*CreateRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationships_ReadRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipsServer).ReadRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relationships_ReadRelationships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipsServer).ReadRelationships(ctx, req.(*ReadRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationships_DeleteRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipsServer).DeleteRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relationships_DeleteRelationships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipsServer).DeleteRelationships(ctx, req.(*DeleteRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Relationships_ServiceDesc is the grpc.ServiceDesc for Relationships service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relationships_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rebac.v1.Relationships",
	HandlerType: (*RelationshipsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRelationships",
			Handler:    _Relationships_CreateRelationships_Handler,
		},
		{
			MethodName: "ReadRelationships",
			Handler:    _Relationships_ReadRelationships_Handler,
		},
		{
			MethodName: "DeleteRelationships",
			Handler:    _Relationships_DeleteRelationships_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rebac/v1/relationships.proto",
}
