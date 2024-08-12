// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Users_Register_FullMethodName       = "/user.Users/Register"
	Users_Login_FullMethodName          = "/user.Users/Login"
	Users_GetUserProfile_FullMethodName = "/user.Users/GetUserProfile"
	Users_UpdateProfile_FullMethodName  = "/user.Users/UpdateProfile"
	Users_GetUSerByEmail_FullMethodName = "/user.Users/GetUSerByEmail"
	Users_GetUsers_FullMethodName       = "/user.Users/GetUsers"
	Users_UpdateUser_FullMethodName     = "/user.Users/UpdateUser"
	Users_DeleteUser_FullMethodName     = "/user.Users/DeleteUser"
	Users_GetUserById_FullMethodName    = "/user.Users/GetUserById"
	Users_UpdatePassword_FullMethodName = "/user.Users/UpdatePassword"
	Users_ChangePassword_FullMethodName = "/user.Users/ChangePassword"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	GetUserProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileRes, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*Void, error)
	GetUSerByEmail(ctx context.Context, in *GetUSerByEmailReq, opts ...grpc.CallOption) (*FilterUsers, error)
	GetUsers(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Void, error)
	GetUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*Void, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordRes, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, Users_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, Users_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileRes)
	err := c.cc.Invoke(ctx, Users_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, Users_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUSerByEmail(ctx context.Context, in *GetUSerByEmailReq, opts ...grpc.CallOption) (*FilterUsers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FilterUsers)
	err := c.cc.Invoke(ctx, Users_GetUSerByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUsers(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, Users_GetUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, Users_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, Users_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, Users_GetUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, Users_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePasswordRes)
	err := c.cc.Invoke(ctx, Users_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	Login(context.Context, *LoginReq) (*LoginRes, error)
	GetUserProfile(context.Context, *GetProfileReq) (*GetProfileRes, error)
	UpdateProfile(context.Context, *UpdateProfileReq) (*Void, error)
	GetUSerByEmail(context.Context, *GetUSerByEmailReq) (*FilterUsers, error)
	GetUsers(context.Context, *UsersListRequest) (*UsersResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*Void, error)
	DeleteUser(context.Context, *UserId) (*Void, error)
	GetUserById(context.Context, *UserId) (*User, error)
	UpdatePassword(context.Context, *UpdatePasswordReq) (*Void, error)
	ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordRes, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) Register(context.Context, *RegisterReq) (*RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUsersServer) Login(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsersServer) GetUserProfile(context.Context, *GetProfileReq) (*GetProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUsersServer) UpdateProfile(context.Context, *UpdateProfileReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUsersServer) GetUSerByEmail(context.Context, *GetUSerByEmailReq) (*FilterUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUSerByEmail not implemented")
}
func (UnimplementedUsersServer) GetUsers(context.Context, *UsersListRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUsersServer) UpdateUser(context.Context, *UpdateUserRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsersServer) DeleteUser(context.Context, *UserId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUsersServer) GetUserById(context.Context, *UserId) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUsersServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUsersServer) ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserProfile(ctx, req.(*GetProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateProfile(ctx, req.(*UpdateProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUSerByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUSerByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUSerByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUSerByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUSerByEmail(ctx, req.(*GetUSerByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUsers(ctx, req.(*UsersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserById(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangePassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Users_Login_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _Users_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Users_UpdateProfile_Handler,
		},
		{
			MethodName: "GetUSerByEmail",
			Handler:    _Users_GetUSerByEmail_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Users_GetUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Users_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Users_GetUserById_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Users_UpdatePassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Users_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
