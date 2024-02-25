// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/auth.proto

package api

import (
	context "context"
	model "github.com/karamaru-alpha/melt/pkg/domain/proto/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Signup(ctx context.Context, in *AuthSignupRequest, opts ...grpc.CallOption) (*AuthSignupResponse, error)
	RefreshToken(ctx context.Context, in *AuthRefreshTokenRequest, opts ...grpc.CallOption) (*AuthRefreshTokenResponse, error)
	Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.User, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Signup(ctx context.Context, in *AuthSignupRequest, opts ...grpc.CallOption) (*AuthSignupResponse, error) {
	out := new(AuthSignupResponse)
	err := c.cc.Invoke(ctx, "/api.Auth/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshToken(ctx context.Context, in *AuthRefreshTokenRequest, opts ...grpc.CallOption) (*AuthRefreshTokenResponse, error) {
	out := new(AuthRefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/api.Auth/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.User, error) {
	out := new(model.User)
	err := c.cc.Invoke(ctx, "/api.Auth/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Signup(context.Context, *AuthSignupRequest) (*AuthSignupResponse, error)
	RefreshToken(context.Context, *AuthRefreshTokenRequest) (*AuthRefreshTokenResponse, error)
	Get(context.Context, *emptypb.Empty) (*model.User, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Signup(context.Context, *AuthSignupRequest) (*AuthSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedAuthServer) RefreshToken(context.Context, *AuthRefreshTokenRequest) (*AuthRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServer) Get(context.Context, *emptypb.Empty) (*model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Auth/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Signup(ctx, req.(*AuthSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Auth/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshToken(ctx, req.(*AuthRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Auth/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Get(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _Auth_Signup_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Auth_RefreshToken_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Auth_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/auth.proto",
}
