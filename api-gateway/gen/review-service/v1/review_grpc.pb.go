// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewServiceClient interface {
	ReviewHealthCheck(ctx context.Context, in *ReviewHealthCheckRequest, opts ...grpc.CallOption) (*ReviewHealthCheckResponse, error)
	CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error)
	ListReviewsByCompany(ctx context.Context, in *ListReviewsByCompanyRequest, opts ...grpc.CallOption) (*ListReviewsByCompanyResponse, error)
	ListReviewsByUser(ctx context.Context, in *ListReviewsByUserRequest, opts ...grpc.CallOption) (*ListReviewsByUserResponse, error)
	GetReview(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*GetReviewResponse, error)
	UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error)
	DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewResponse, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) ReviewHealthCheck(ctx context.Context, in *ReviewHealthCheckRequest, opts ...grpc.CallOption) (*ReviewHealthCheckResponse, error) {
	out := new(ReviewHealthCheckResponse)
	err := c.cc.Invoke(ctx, "/user.ReviewService/ReviewHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error) {
	out := new(CreateReviewResponse)
	err := c.cc.Invoke(ctx, "/user.ReviewService/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) ListReviewsByCompany(ctx context.Context, in *ListReviewsByCompanyRequest, opts ...grpc.CallOption) (*ListReviewsByCompanyResponse, error) {
	out := new(ListReviewsByCompanyResponse)
	err := c.cc.Invoke(ctx, "/user.ReviewService/ListReviewsByCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) ListReviewsByUser(ctx context.Context, in *ListReviewsByUserRequest, opts ...grpc.CallOption) (*ListReviewsByUserResponse, error) {
	out := new(ListReviewsByUserResponse)
	err := c.cc.Invoke(ctx, "/user.ReviewService/ListReviewsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReview(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*GetReviewResponse, error) {
	out := new(GetReviewResponse)
	err := c.cc.Invoke(ctx, "/user.ReviewService/GetReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error) {
	out := new(UpdateReviewResponse)
	err := c.cc.Invoke(ctx, "/user.ReviewService/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewResponse, error) {
	out := new(DeleteReviewResponse)
	err := c.cc.Invoke(ctx, "/user.ReviewService/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
// All implementations must embed UnimplementedReviewServiceServer
// for forward compatibility
type ReviewServiceServer interface {
	ReviewHealthCheck(context.Context, *ReviewHealthCheckRequest) (*ReviewHealthCheckResponse, error)
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error)
	ListReviewsByCompany(context.Context, *ListReviewsByCompanyRequest) (*ListReviewsByCompanyResponse, error)
	ListReviewsByUser(context.Context, *ListReviewsByUserRequest) (*ListReviewsByUserResponse, error)
	GetReview(context.Context, *GetReviewRequest) (*GetReviewResponse, error)
	UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error)
	DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewResponse, error)
	mustEmbedUnimplementedReviewServiceServer()
}

// UnimplementedReviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (UnimplementedReviewServiceServer) ReviewHealthCheck(context.Context, *ReviewHealthCheckRequest) (*ReviewHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewHealthCheck not implemented")
}
func (UnimplementedReviewServiceServer) CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedReviewServiceServer) ListReviewsByCompany(context.Context, *ListReviewsByCompanyRequest) (*ListReviewsByCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReviewsByCompany not implemented")
}
func (UnimplementedReviewServiceServer) ListReviewsByUser(context.Context, *ListReviewsByUserRequest) (*ListReviewsByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReviewsByUser not implemented")
}
func (UnimplementedReviewServiceServer) GetReview(context.Context, *GetReviewRequest) (*GetReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReview not implemented")
}
func (UnimplementedReviewServiceServer) UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedReviewServiceServer) DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {}

// UnsafeReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServiceServer will
// result in compilation errors.
type UnsafeReviewServiceServer interface {
	mustEmbedUnimplementedReviewServiceServer()
}

func RegisterReviewServiceServer(s grpc.ServiceRegistrar, srv ReviewServiceServer) {
	s.RegisterService(&ReviewService_ServiceDesc, srv)
}

func _ReviewService_ReviewHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ReviewHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReviewService/ReviewHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ReviewHealthCheck(ctx, req.(*ReviewHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReviewService/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_ListReviewsByCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReviewsByCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ListReviewsByCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReviewService/ListReviewsByCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ListReviewsByCompany(ctx, req.(*ListReviewsByCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_ListReviewsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReviewsByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ListReviewsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReviewService/ListReviewsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ListReviewsByUser(ctx, req.(*ListReviewsByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReviewService/GetReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReview(ctx, req.(*GetReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReviewService/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).UpdateReview(ctx, req.(*UpdateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReviewService/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).DeleteReview(ctx, req.(*DeleteReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewService_ServiceDesc is the grpc.ServiceDesc for ReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReviewHealthCheck",
			Handler:    _ReviewService_ReviewHealthCheck_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _ReviewService_CreateReview_Handler,
		},
		{
			MethodName: "ListReviewsByCompany",
			Handler:    _ReviewService_ListReviewsByCompany_Handler,
		},
		{
			MethodName: "ListReviewsByUser",
			Handler:    _ReviewService_ListReviewsByUser_Handler,
		},
		{
			MethodName: "GetReview",
			Handler:    _ReviewService_GetReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _ReviewService_UpdateReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _ReviewService_DeleteReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/review.proto",
}
