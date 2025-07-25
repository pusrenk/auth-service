// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: main.proto

package protogen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Main_GetUserBySessionID_FullMethodName = "/main.Main/GetUserBySessionID"
	Main_StoreUserSession_FullMethodName   = "/main.Main/StoreUserSession"
)

// MainClient is the client API for Main service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainClient interface {
	GetUserBySessionID(ctx context.Context, in *GetUserBySessionIDRequest, opts ...grpc.CallOption) (*UserResponse, error)
	StoreUserSession(ctx context.Context, in *StoreUserSessionRequest, opts ...grpc.CallOption) (*Empty, error)
}

type mainClient struct {
	cc grpc.ClientConnInterface
}

func NewMainClient(cc grpc.ClientConnInterface) MainClient {
	return &mainClient{cc}
}

func (c *mainClient) GetUserBySessionID(ctx context.Context, in *GetUserBySessionIDRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, Main_GetUserBySessionID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) StoreUserSession(ctx context.Context, in *StoreUserSessionRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_StoreUserSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServer is the server API for Main service.
// All implementations must embed UnimplementedMainServer
// for forward compatibility.
type MainServer interface {
	GetUserBySessionID(context.Context, *GetUserBySessionIDRequest) (*UserResponse, error)
	StoreUserSession(context.Context, *StoreUserSessionRequest) (*Empty, error)
	mustEmbedUnimplementedMainServer()
}

// UnimplementedMainServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMainServer struct{}

func (UnimplementedMainServer) GetUserBySessionID(context.Context, *GetUserBySessionIDRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBySessionID not implemented")
}
func (UnimplementedMainServer) StoreUserSession(context.Context, *StoreUserSessionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreUserSession not implemented")
}
func (UnimplementedMainServer) mustEmbedUnimplementedMainServer() {}
func (UnimplementedMainServer) testEmbeddedByValue()              {}

// UnsafeMainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServer will
// result in compilation errors.
type UnsafeMainServer interface {
	mustEmbedUnimplementedMainServer()
}

func RegisterMainServer(s grpc.ServiceRegistrar, srv MainServer) {
	// If the following call pancis, it indicates UnimplementedMainServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Main_ServiceDesc, srv)
}

func _Main_GetUserBySessionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBySessionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).GetUserBySessionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_GetUserBySessionID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).GetUserBySessionID(ctx, req.(*GetUserBySessionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_StoreUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreUserSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).StoreUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_StoreUserSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).StoreUserSession(ctx, req.(*StoreUserSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Main_ServiceDesc is the grpc.ServiceDesc for Main service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Main_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Main",
	HandlerType: (*MainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserBySessionID",
			Handler:    _Main_GetUserBySessionID_Handler,
		},
		{
			MethodName: "StoreUserSession",
			Handler:    _Main_StoreUserSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
