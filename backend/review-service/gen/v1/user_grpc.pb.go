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

const (
	UserService_UserHealthCheck_FullMethodName       = "/user.UserService/UserHealthCheck"
	UserService_GetStudentMe_FullMethodName          = "/user.UserService/GetStudentMe"
	UserService_GetStudent_FullMethodName            = "/user.UserService/GetStudent"
	UserService_UpdateStudent_FullMethodName         = "/user.UserService/UpdateStudent"
	UserService_GetCompanyMe_FullMethodName          = "/user.UserService/GetCompanyMe"
	UserService_GetCompany_FullMethodName            = "/user.UserService/GetCompany"
	UserService_UpdateCompany_FullMethodName         = "/user.UserService/UpdateCompany"
	UserService_ListApprovedCompanies_FullMethodName = "/user.UserService/ListApprovedCompanies"
	UserService_ListCompanies_FullMethodName         = "/user.UserService/ListCompanies"
	UserService_UpdateCompanyStatus_FullMethodName   = "/user.UserService/UpdateCompanyStatus"
	UserService_GetStudents_FullMethodName           = "/user.UserService/GetStudents"
	UserService_GetCompanies_FullMethodName          = "/user.UserService/GetCompanies"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserHealthCheck(ctx context.Context, in *UserHealthCheckRequest, opts ...grpc.CallOption) (*UserHealthCheckResponse, error)
	GetStudentMe(ctx context.Context, in *GetStudentMeRequest, opts ...grpc.CallOption) (*GetStudentResponse, error)
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error)
	UpdateStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error)
	GetCompanyMe(ctx context.Context, in *GetCompanyMeRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
	ListApprovedCompanies(ctx context.Context, in *ListApprovedCompaniesRequest, opts ...grpc.CallOption) (*ListApprovedCompaniesResponse, error)
	ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error)
	UpdateCompanyStatus(ctx context.Context, in *UpdateCompanyStatusRequest, opts ...grpc.CallOption) (*UpdateCompanyStatusResponse, error)
	GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error)
	GetCompanies(ctx context.Context, in *GetCompaniesRequest, opts ...grpc.CallOption) (*GetCompaniesResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserHealthCheck(ctx context.Context, in *UserHealthCheckRequest, opts ...grpc.CallOption) (*UserHealthCheckResponse, error) {
	out := new(UserHealthCheckResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UserHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetStudentMe(ctx context.Context, in *GetStudentMeRequest, opts ...grpc.CallOption) (*GetStudentResponse, error) {
	out := new(GetStudentResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetStudentMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error) {
	out := new(GetStudentResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error) {
	out := new(UpdateStudentResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCompanyMe(ctx context.Context, in *GetCompanyMeRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetCompanyMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListApprovedCompanies(ctx context.Context, in *ListApprovedCompaniesRequest, opts ...grpc.CallOption) (*ListApprovedCompaniesResponse, error) {
	out := new(ListApprovedCompaniesResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ListApprovedCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error) {
	out := new(ListCompaniesResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ListCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCompanyStatus(ctx context.Context, in *UpdateCompanyStatusRequest, opts ...grpc.CallOption) (*UpdateCompanyStatusResponse, error) {
	out := new(UpdateCompanyStatusResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateCompanyStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error) {
	out := new(GetStudentsResponse)
	err := c.cc.Invoke(ctx, UserService_GetStudents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCompanies(ctx context.Context, in *GetCompaniesRequest, opts ...grpc.CallOption) (*GetCompaniesResponse, error) {
	out := new(GetCompaniesResponse)
	err := c.cc.Invoke(ctx, UserService_GetCompanies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserHealthCheck(context.Context, *UserHealthCheckRequest) (*UserHealthCheckResponse, error)
	GetStudentMe(context.Context, *GetStudentMeRequest) (*GetStudentResponse, error)
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error)
	UpdateStudent(context.Context, *UpdateStudentRequest) (*UpdateStudentResponse, error)
	GetCompanyMe(context.Context, *GetCompanyMeRequest) (*GetCompanyResponse, error)
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	ListApprovedCompanies(context.Context, *ListApprovedCompaniesRequest) (*ListApprovedCompaniesResponse, error)
	ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error)
	UpdateCompanyStatus(context.Context, *UpdateCompanyStatusRequest) (*UpdateCompanyStatusResponse, error)
	GetStudents(context.Context, *GetStudentsRequest) (*GetStudentsResponse, error)
	GetCompanies(context.Context, *GetCompaniesRequest) (*GetCompaniesResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserHealthCheck(context.Context, *UserHealthCheckRequest) (*UserHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserHealthCheck not implemented")
}
func (UnimplementedUserServiceServer) GetStudentMe(context.Context, *GetStudentMeRequest) (*GetStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentMe not implemented")
}
func (UnimplementedUserServiceServer) GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedUserServiceServer) UpdateStudent(context.Context, *UpdateStudentRequest) (*UpdateStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedUserServiceServer) GetCompanyMe(context.Context, *GetCompanyMeRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyMe not implemented")
}
func (UnimplementedUserServiceServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedUserServiceServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedUserServiceServer) ListApprovedCompanies(context.Context, *ListApprovedCompaniesRequest) (*ListApprovedCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApprovedCompanies not implemented")
}
func (UnimplementedUserServiceServer) ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanies not implemented")
}
func (UnimplementedUserServiceServer) UpdateCompanyStatus(context.Context, *UpdateCompanyStatusRequest) (*UpdateCompanyStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompanyStatus not implemented")
}
func (UnimplementedUserServiceServer) GetStudents(context.Context, *GetStudentsRequest) (*GetStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}
func (UnimplementedUserServiceServer) GetCompanies(context.Context, *GetCompaniesRequest) (*GetCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanies not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserHealthCheck(ctx, req.(*UserHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetStudentMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetStudentMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetStudentMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetStudentMe(ctx, req.(*GetStudentMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateStudent(ctx, req.(*UpdateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCompanyMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCompanyMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetCompanyMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCompanyMe(ctx, req.(*GetCompanyMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListApprovedCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApprovedCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListApprovedCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListApprovedCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListApprovedCompanies(ctx, req.(*ListApprovedCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListCompanies(ctx, req.(*ListCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCompanyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCompanyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateCompanyStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCompanyStatus(ctx, req.(*UpdateCompanyStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetStudents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetStudents(ctx, req.(*GetStudentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCompanies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCompanies(ctx, req.(*GetCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserHealthCheck",
			Handler:    _UserService_UserHealthCheck_Handler,
		},
		{
			MethodName: "GetStudentMe",
			Handler:    _UserService_GetStudentMe_Handler,
		},
		{
			MethodName: "GetStudent",
			Handler:    _UserService_GetStudent_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _UserService_UpdateStudent_Handler,
		},
		{
			MethodName: "GetCompanyMe",
			Handler:    _UserService_GetCompanyMe_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _UserService_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _UserService_UpdateCompany_Handler,
		},
		{
			MethodName: "ListApprovedCompanies",
			Handler:    _UserService_ListApprovedCompanies_Handler,
		},
		{
			MethodName: "ListCompanies",
			Handler:    _UserService_ListCompanies_Handler,
		},
		{
			MethodName: "UpdateCompanyStatus",
			Handler:    _UserService_UpdateCompanyStatus_Handler,
		},
		{
			MethodName: "GetStudents",
			Handler:    _UserService_GetStudents_Handler,
		},
		{
			MethodName: "GetCompanies",
			Handler:    _UserService_GetCompanies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/user.proto",
}