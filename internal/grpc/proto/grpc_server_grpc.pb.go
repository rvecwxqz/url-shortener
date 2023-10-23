// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: internal/grpc/proto/grpc_server.proto

package proto

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

const (
	Shortener_GetLongUrl_FullMethodName  = "/grpc.Shortener/GetLongUrl"
	Shortener_GetShortUrl_FullMethodName = "/grpc.Shortener/GetShortUrl"
	Shortener_GetUrls_FullMethodName     = "/grpc.Shortener/GetUrls"
	Shortener_DeleteUrls_FullMethodName  = "/grpc.Shortener/DeleteUrls"
	Shortener_PingDb_FullMethodName      = "/grpc.Shortener/PingDb"
)

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerClient interface {
	GetLongUrl(ctx context.Context, in *GetLongUrlRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error)
	GetUrls(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUrlsResponse, error)
	DeleteUrls(ctx context.Context, in *DeleteUrlsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PingDb(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) GetLongUrl(ctx context.Context, in *GetLongUrlRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Shortener_GetLongUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error) {
	out := new(GetShortUrlResponse)
	err := c.cc.Invoke(ctx, Shortener_GetShortUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetUrls(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUrlsResponse, error) {
	out := new(GetUrlsResponse)
	err := c.cc.Invoke(ctx, Shortener_GetUrls_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) DeleteUrls(ctx context.Context, in *DeleteUrlsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Shortener_DeleteUrls_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) PingDb(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Shortener_PingDb_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
// All implementations must embed UnimplementedShortenerServer
// for forward compatibility
type ShortenerServer interface {
	GetLongUrl(context.Context, *GetLongUrlRequest) (*emptypb.Empty, error)
	GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error)
	GetUrls(context.Context, *emptypb.Empty) (*GetUrlsResponse, error)
	DeleteUrls(context.Context, *DeleteUrlsRequest) (*emptypb.Empty, error)
	PingDb(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedShortenerServer()
}

// UnimplementedShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedShortenerServer struct {
}

func (UnimplementedShortenerServer) GetLongUrl(context.Context, *GetLongUrlRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLongUrl not implemented")
}
func (UnimplementedShortenerServer) GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortUrl not implemented")
}
func (UnimplementedShortenerServer) GetUrls(context.Context, *emptypb.Empty) (*GetUrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrls not implemented")
}
func (UnimplementedShortenerServer) DeleteUrls(context.Context, *DeleteUrlsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUrls not implemented")
}
func (UnimplementedShortenerServer) PingDb(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingDb not implemented")
}
func (UnimplementedShortenerServer) mustEmbedUnimplementedShortenerServer() {}

// UnsafeShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServer will
// result in compilation errors.
type UnsafeShortenerServer interface {
	mustEmbedUnimplementedShortenerServer()
}

func RegisterShortenerServer(s grpc.ServiceRegistrar, srv ShortenerServer) {
	s.RegisterService(&Shortener_ServiceDesc, srv)
}

func _Shortener_GetLongUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetLongUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_GetLongUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetLongUrl(ctx, req.(*GetLongUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_GetShortUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetShortUrl(ctx, req.(*GetShortUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_GetUrls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetUrls(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_DeleteUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).DeleteUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_DeleteUrls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).DeleteUrls(ctx, req.(*DeleteUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_PingDb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).PingDb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_PingDb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).PingDb(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Shortener_ServiceDesc is the grpc.ServiceDesc for Shortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLongUrl",
			Handler:    _Shortener_GetLongUrl_Handler,
		},
		{
			MethodName: "GetShortUrl",
			Handler:    _Shortener_GetShortUrl_Handler,
		},
		{
			MethodName: "GetUrls",
			Handler:    _Shortener_GetUrls_Handler,
		},
		{
			MethodName: "DeleteUrls",
			Handler:    _Shortener_DeleteUrls_Handler,
		},
		{
			MethodName: "PingDb",
			Handler:    _Shortener_PingDb_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/grpc/proto/grpc_server.proto",
}
