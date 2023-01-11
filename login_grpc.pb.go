// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: login.proto

package puzzleloginservice

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

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	Verify(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	Register(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	ListLogins(ctx context.Context, in *UserIds, opts ...grpc.CallOption) (*Logins, error)
	ChangeLogin(ctx context.Context, in *ChangeLoginRequest, opts ...grpc.CallOption) (*Response, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*Response, error)
	ListUsers(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*Users, error)
	Delete(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Response, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Verify(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/puzzleloginservice.Login/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) Register(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/puzzleloginservice.Login/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) ListLogins(ctx context.Context, in *UserIds, opts ...grpc.CallOption) (*Logins, error) {
	out := new(Logins)
	err := c.cc.Invoke(ctx, "/puzzleloginservice.Login/ListLogins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) ChangeLogin(ctx context.Context, in *ChangeLoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/puzzleloginservice.Login/ChangeLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/puzzleloginservice.Login/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) ListUsers(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/puzzleloginservice.Login/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) Delete(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/puzzleloginservice.Login/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	Verify(context.Context, *LoginRequest) (*Response, error)
	Register(context.Context, *LoginRequest) (*Response, error)
	ListLogins(context.Context, *UserIds) (*Logins, error)
	ChangeLogin(context.Context, *ChangeLoginRequest) (*Response, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*Response, error)
	ListUsers(context.Context, *RangeRequest) (*Users, error)
	Delete(context.Context, *UserId) (*Response, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) Verify(context.Context, *LoginRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedLoginServer) Register(context.Context, *LoginRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedLoginServer) ListLogins(context.Context, *UserIds) (*Logins, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogins not implemented")
}
func (UnimplementedLoginServer) ChangeLogin(context.Context, *ChangeLoginRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeLogin not implemented")
}
func (UnimplementedLoginServer) ChangePassword(context.Context, *ChangePasswordRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedLoginServer) ListUsers(context.Context, *RangeRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedLoginServer) Delete(context.Context, *UserId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzleloginservice.Login/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Verify(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzleloginservice.Login/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Register(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_ListLogins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).ListLogins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzleloginservice.Login/ListLogins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).ListLogins(ctx, req.(*UserIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_ChangeLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).ChangeLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzleloginservice.Login/ChangeLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).ChangeLogin(ctx, req.(*ChangeLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzleloginservice.Login/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzleloginservice.Login/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).ListUsers(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzleloginservice.Login/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Delete(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc.ServiceDesc for Login service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "puzzleloginservice.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _Login_Verify_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Login_Register_Handler,
		},
		{
			MethodName: "ListLogins",
			Handler:    _Login_ListLogins_Handler,
		},
		{
			MethodName: "ChangeLogin",
			Handler:    _Login_ChangeLogin_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Login_ChangePassword_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Login_ListUsers_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Login_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
