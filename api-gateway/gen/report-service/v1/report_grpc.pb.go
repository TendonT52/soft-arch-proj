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

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	ReportHealthCheck(ctx context.Context, in *ReportHealthCheckRequest, opts ...grpc.CallOption) (*ReportHealthCheckResponse, error)
	CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*CreateReportResponse, error)
	ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error)
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) ReportHealthCheck(ctx context.Context, in *ReportHealthCheckRequest, opts ...grpc.CallOption) (*ReportHealthCheckResponse, error) {
	out := new(ReportHealthCheckResponse)
	err := c.cc.Invoke(ctx, "/user.ReportService/ReportHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*CreateReportResponse, error) {
	out := new(CreateReportResponse)
	err := c.cc.Invoke(ctx, "/user.ReportService/CreateReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error) {
	out := new(ListReportsResponse)
	err := c.cc.Invoke(ctx, "/user.ReportService/ListReports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error) {
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, "/user.ReportService/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations must embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	ReportHealthCheck(context.Context, *ReportHealthCheckRequest) (*ReportHealthCheckResponse, error)
	CreateReport(context.Context, *CreateReportRequest) (*CreateReportResponse, error)
	ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error)
	GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error)
	mustEmbedUnimplementedReportServiceServer()
}

// UnimplementedReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (UnimplementedReportServiceServer) ReportHealthCheck(context.Context, *ReportHealthCheckRequest) (*ReportHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportHealthCheck not implemented")
}
func (UnimplementedReportServiceServer) CreateReport(context.Context, *CreateReportRequest) (*CreateReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReport not implemented")
}
func (UnimplementedReportServiceServer) ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReports not implemented")
}
func (UnimplementedReportServiceServer) GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedReportServiceServer) mustEmbedUnimplementedReportServiceServer() {}

// UnsafeReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServiceServer will
// result in compilation errors.
type UnsafeReportServiceServer interface {
	mustEmbedUnimplementedReportServiceServer()
}

func RegisterReportServiceServer(s grpc.ServiceRegistrar, srv ReportServiceServer) {
	s.RegisterService(&ReportService_ServiceDesc, srv)
}

func _ReportService_ReportHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).ReportHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReportService/ReportHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).ReportHealthCheck(ctx, req.(*ReportHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_CreateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).CreateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReportService/CreateReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).CreateReport(ctx, req.(*CreateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_ListReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).ListReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReportService/ListReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).ListReports(ctx, req.(*ListReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ReportService/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportService_ServiceDesc is the grpc.ServiceDesc for ReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportHealthCheck",
			Handler:    _ReportService_ReportHealthCheck_Handler,
		},
		{
			MethodName: "CreateReport",
			Handler:    _ReportService_CreateReport_Handler,
		},
		{
			MethodName: "ListReports",
			Handler:    _ReportService_ListReports_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _ReportService_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/report.proto",
}