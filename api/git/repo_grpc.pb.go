// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: api/git/repo.proto

package git

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
	Repo_CreateRepo_FullMethodName = "/api.git.Repo/CreateRepo"
	Repo_UpdateRepo_FullMethodName = "/api.git.Repo/UpdateRepo"
	Repo_DeleteRepo_FullMethodName = "/api.git.Repo/DeleteRepo"
	Repo_GetRepo_FullMethodName    = "/api.git.Repo/GetRepo"
	Repo_ListRepo_FullMethodName   = "/api.git.Repo/ListRepo"
)

// RepoClient is the client API for Repo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepoClient interface {
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoReply, error)
	UpdateRepo(ctx context.Context, in *UpdateRepoRequest, opts ...grpc.CallOption) (*UpdateRepoReply, error)
	DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoReply, error)
	GetRepo(ctx context.Context, in *GetRepoRequest, opts ...grpc.CallOption) (*GetRepoReply, error)
	ListRepo(ctx context.Context, in *ListRepoRequest, opts ...grpc.CallOption) (*ListRepoReply, error)
}

type repoClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoClient(cc grpc.ClientConnInterface) RepoClient {
	return &repoClient{cc}
}

func (c *repoClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoReply, error) {
	out := new(CreateRepoReply)
	err := c.cc.Invoke(ctx, Repo_CreateRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) UpdateRepo(ctx context.Context, in *UpdateRepoRequest, opts ...grpc.CallOption) (*UpdateRepoReply, error) {
	out := new(UpdateRepoReply)
	err := c.cc.Invoke(ctx, Repo_UpdateRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoReply, error) {
	out := new(DeleteRepoReply)
	err := c.cc.Invoke(ctx, Repo_DeleteRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) GetRepo(ctx context.Context, in *GetRepoRequest, opts ...grpc.CallOption) (*GetRepoReply, error) {
	out := new(GetRepoReply)
	err := c.cc.Invoke(ctx, Repo_GetRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) ListRepo(ctx context.Context, in *ListRepoRequest, opts ...grpc.CallOption) (*ListRepoReply, error) {
	out := new(ListRepoReply)
	err := c.cc.Invoke(ctx, Repo_ListRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoServer is the server API for Repo service.
// All implementations must embed UnimplementedRepoServer
// for forward compatibility
type RepoServer interface {
	CreateRepo(context.Context, *CreateRepoRequest) (*CreateRepoReply, error)
	UpdateRepo(context.Context, *UpdateRepoRequest) (*UpdateRepoReply, error)
	DeleteRepo(context.Context, *DeleteRepoRequest) (*DeleteRepoReply, error)
	GetRepo(context.Context, *GetRepoRequest) (*GetRepoReply, error)
	ListRepo(context.Context, *ListRepoRequest) (*ListRepoReply, error)
	mustEmbedUnimplementedRepoServer()
}

// UnimplementedRepoServer must be embedded to have forward compatible implementations.
type UnimplementedRepoServer struct {
}

func (UnimplementedRepoServer) CreateRepo(context.Context, *CreateRepoRequest) (*CreateRepoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepo not implemented")
}
func (UnimplementedRepoServer) UpdateRepo(context.Context, *UpdateRepoRequest) (*UpdateRepoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepo not implemented")
}
func (UnimplementedRepoServer) DeleteRepo(context.Context, *DeleteRepoRequest) (*DeleteRepoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepo not implemented")
}
func (UnimplementedRepoServer) GetRepo(context.Context, *GetRepoRequest) (*GetRepoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepo not implemented")
}
func (UnimplementedRepoServer) ListRepo(context.Context, *ListRepoRequest) (*ListRepoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepo not implemented")
}
func (UnimplementedRepoServer) mustEmbedUnimplementedRepoServer() {}

// UnsafeRepoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepoServer will
// result in compilation errors.
type UnsafeRepoServer interface {
	mustEmbedUnimplementedRepoServer()
}

func RegisterRepoServer(s grpc.ServiceRegistrar, srv RepoServer) {
	s.RegisterService(&Repo_ServiceDesc, srv)
}

func _Repo_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_CreateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_UpdateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).UpdateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_UpdateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).UpdateRepo(ctx, req.(*UpdateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_DeleteRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).DeleteRepo(ctx, req.(*DeleteRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_GetRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).GetRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_GetRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).GetRepo(ctx, req.(*GetRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_ListRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).ListRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_ListRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).ListRepo(ctx, req.(*ListRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Repo_ServiceDesc is the grpc.ServiceDesc for Repo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.git.Repo",
	HandlerType: (*RepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepo",
			Handler:    _Repo_CreateRepo_Handler,
		},
		{
			MethodName: "UpdateRepo",
			Handler:    _Repo_UpdateRepo_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _Repo_DeleteRepo_Handler,
		},
		{
			MethodName: "GetRepo",
			Handler:    _Repo_GetRepo_Handler,
		},
		{
			MethodName: "ListRepo",
			Handler:    _Repo_ListRepo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/git/repo.proto",
}
