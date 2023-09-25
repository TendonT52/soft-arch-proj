// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: v1/post.proto

package gen

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
	PostService_PostHealthCheck_FullMethodName   = "/user.PostService/PostHealthCheck"
	PostService_CreatePost_FullMethodName        = "/user.PostService/CreatePost"
	PostService_GetPost_FullMethodName           = "/user.PostService/GetPost"
	PostService_ListPosts_FullMethodName         = "/user.PostService/ListPosts"
	PostService_UpdatePost_FullMethodName        = "/user.PostService/UpdatePost"
	PostService_DeletePost_FullMethodName        = "/user.PostService/DeletePost"
	PostService_DeletePosts_FullMethodName       = "/user.PostService/DeletePosts"
	PostService_GetOpenPositions_FullMethodName  = "/user.PostService/GetOpenPositions"
	PostService_GetRequiredSkills_FullMethodName = "/user.PostService/GetRequiredSkills"
	PostService_GetBenefits_FullMethodName       = "/user.PostService/GetBenefits"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	PostHealthCheck(ctx context.Context, in *PostHealthCheckRequest, opts ...grpc.CallOption) (*PostHealthCheckResponse, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	DeletePosts(ctx context.Context, in *DeletePostsRequest, opts ...grpc.CallOption) (*DeletePostsResponse, error)
	GetOpenPositions(ctx context.Context, in *GetOpenPositionsRequest, opts ...grpc.CallOption) (*GetOpenPositionsResponse, error)
	GetRequiredSkills(ctx context.Context, in *GetRequiredSkillsRequest, opts ...grpc.CallOption) (*GetRequiredSkillsResponse, error)
	GetBenefits(ctx context.Context, in *GetBenefitsRequest, opts ...grpc.CallOption) (*GetBenefitsResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) PostHealthCheck(ctx context.Context, in *PostHealthCheckRequest, opts ...grpc.CallOption) (*PostHealthCheckResponse, error) {
	out := new(PostHealthCheckResponse)
	err := c.cc.Invoke(ctx, PostService_PostHealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, PostService_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error) {
	out := new(ListPostsResponse)
	err := c.cc.Invoke(ctx, PostService_ListPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, PostService_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, PostService_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePosts(ctx context.Context, in *DeletePostsRequest, opts ...grpc.CallOption) (*DeletePostsResponse, error) {
	out := new(DeletePostsResponse)
	err := c.cc.Invoke(ctx, PostService_DeletePosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetOpenPositions(ctx context.Context, in *GetOpenPositionsRequest, opts ...grpc.CallOption) (*GetOpenPositionsResponse, error) {
	out := new(GetOpenPositionsResponse)
	err := c.cc.Invoke(ctx, PostService_GetOpenPositions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetRequiredSkills(ctx context.Context, in *GetRequiredSkillsRequest, opts ...grpc.CallOption) (*GetRequiredSkillsResponse, error) {
	out := new(GetRequiredSkillsResponse)
	err := c.cc.Invoke(ctx, PostService_GetRequiredSkills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetBenefits(ctx context.Context, in *GetBenefitsRequest, opts ...grpc.CallOption) (*GetBenefitsResponse, error) {
	out := new(GetBenefitsResponse)
	err := c.cc.Invoke(ctx, PostService_GetBenefits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	PostHealthCheck(context.Context, *PostHealthCheckRequest) (*PostHealthCheckResponse, error)
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	DeletePosts(context.Context, *DeletePostsRequest) (*DeletePostsResponse, error)
	GetOpenPositions(context.Context, *GetOpenPositionsRequest) (*GetOpenPositionsResponse, error)
	GetRequiredSkills(context.Context, *GetRequiredSkillsRequest) (*GetRequiredSkillsResponse, error)
	GetBenefits(context.Context, *GetBenefitsRequest) (*GetBenefitsResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) PostHealthCheck(context.Context, *PostHealthCheckRequest) (*PostHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostHealthCheck not implemented")
}
func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedPostServiceServer) UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePosts(context.Context, *DeletePostsRequest) (*DeletePostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosts not implemented")
}
func (UnimplementedPostServiceServer) GetOpenPositions(context.Context, *GetOpenPositionsRequest) (*GetOpenPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenPositions not implemented")
}
func (UnimplementedPostServiceServer) GetRequiredSkills(context.Context, *GetRequiredSkillsRequest) (*GetRequiredSkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequiredSkills not implemented")
}
func (UnimplementedPostServiceServer) GetBenefits(context.Context, *GetBenefitsRequest) (*GetBenefitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBenefits not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_PostHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_PostHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostHealthCheck(ctx, req.(*PostHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ListPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ListPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ListPosts(ctx, req.(*ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePosts(ctx, req.(*DeletePostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetOpenPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOpenPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetOpenPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetOpenPositions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetOpenPositions(ctx, req.(*GetOpenPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetRequiredSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequiredSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetRequiredSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetRequiredSkills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetRequiredSkills(ctx, req.(*GetRequiredSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetBenefits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBenefitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetBenefits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetBenefits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetBenefits(ctx, req.(*GetBenefitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostHealthCheck",
			Handler:    _PostService_PostHealthCheck_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "ListPosts",
			Handler:    _PostService_ListPosts_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "DeletePosts",
			Handler:    _PostService_DeletePosts_Handler,
		},
		{
			MethodName: "GetOpenPositions",
			Handler:    _PostService_GetOpenPositions_Handler,
		},
		{
			MethodName: "GetRequiredSkills",
			Handler:    _PostService_GetRequiredSkills_Handler,
		},
		{
			MethodName: "GetBenefits",
			Handler:    _PostService_GetBenefits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/post.proto",
}
