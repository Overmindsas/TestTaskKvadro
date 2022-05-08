// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: internal/proto/service.proto

package TestTaskKvadro

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

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	FindBook(ctx context.Context, in *Author, opts ...grpc.CallOption) (*BookList, error)
	FindAuthor(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Authorlist, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) FindBook(ctx context.Context, in *Author, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/internal.API/FindBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) FindAuthor(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Authorlist, error) {
	out := new(Authorlist)
	err := c.cc.Invoke(ctx, "/internal.API/FindAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	FindBook(context.Context, *Author) (*BookList, error)
	FindAuthor(context.Context, *Book) (*Authorlist, error)
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) FindBook(context.Context, *Author) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBook not implemented")
}
func (UnimplementedAPIServer) FindAuthor(context.Context, *Book) (*Authorlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAuthor not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_FindBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).FindBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.API/FindBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).FindBook(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_FindAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).FindAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.API/FindAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).FindAuthor(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBook",
			Handler:    _API_FindBook_Handler,
		},
		{
			MethodName: "FindAuthor",
			Handler:    _API_FindAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/service.proto",
}
