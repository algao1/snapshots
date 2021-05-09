// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RepoClient is the client API for Repo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepoClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (Repo_UploadImageClient, error)
	DownloadImage(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Repo_DownloadImageClient, error)
	ListImages(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type repoClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoClient(cc grpc.ClientConnInterface) RepoClient {
	return &repoClient{cc}
}

func (c *repoClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Repo/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.Repo/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (Repo_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Repo_ServiceDesc.Streams[0], "/proto.Repo/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &repoUploadImageClient{stream}
	return x, nil
}

type Repo_UploadImageClient interface {
	Send(*Upload) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type repoUploadImageClient struct {
	grpc.ClientStream
}

func (x *repoUploadImageClient) Send(m *Upload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *repoUploadImageClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repoClient) DownloadImage(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Repo_DownloadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Repo_ServiceDesc.Streams[1], "/proto.Repo/DownloadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &repoDownloadImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Repo_DownloadImageClient interface {
	Recv() (*Download, error)
	grpc.ClientStream
}

type repoDownloadImageClient struct {
	grpc.ClientStream
}

func (x *repoDownloadImageClient) Recv() (*Download, error) {
	m := new(Download)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repoClient) ListImages(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Repo/ListImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoServer is the server API for Repo service.
// All implementations must embed UnimplementedRepoServer
// for forward compatibility
type RepoServer interface {
	Register(context.Context, *RegisterRequest) (*empty.Empty, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UploadImage(Repo_UploadImageServer) error
	DownloadImage(*DownloadRequest, Repo_DownloadImageServer) error
	ListImages(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedRepoServer()
}

// UnimplementedRepoServer must be embedded to have forward compatible implementations.
type UnimplementedRepoServer struct {
}

func (UnimplementedRepoServer) Register(context.Context, *RegisterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRepoServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRepoServer) UploadImage(Repo_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedRepoServer) DownloadImage(*DownloadRequest, Repo_DownloadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadImage not implemented")
}
func (UnimplementedRepoServer) ListImages(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
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

func _Repo_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repo/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repo/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RepoServer).UploadImage(&repoUploadImageServer{stream})
}

type Repo_UploadImageServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*Upload, error)
	grpc.ServerStream
}

type repoUploadImageServer struct {
	grpc.ServerStream
}

func (x *repoUploadImageServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *repoUploadImageServer) Recv() (*Upload, error) {
	m := new(Upload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Repo_DownloadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RepoServer).DownloadImage(m, &repoDownloadImageServer{stream})
}

type Repo_DownloadImageServer interface {
	Send(*Download) error
	grpc.ServerStream
}

type repoDownloadImageServer struct {
	grpc.ServerStream
}

func (x *repoDownloadImageServer) Send(m *Download) error {
	return x.ServerStream.SendMsg(m)
}

func _Repo_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repo/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).ListImages(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Repo_ServiceDesc is the grpc.ServiceDesc for Repo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Repo",
	HandlerType: (*RepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Repo_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Repo_Login_Handler,
		},
		{
			MethodName: "ListImages",
			Handler:    _Repo_ListImages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _Repo_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadImage",
			Handler:       _Repo_DownloadImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/imgrepo.proto",
}