// Code generated by MockGen. DO NOT EDIT.
// Source: port/service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gen "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthServicePort is a mock of AuthServicePort interface.
type MockAuthServicePort struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServicePortMockRecorder
}

// MockAuthServicePortMockRecorder is the mock recorder for MockAuthServicePort.
type MockAuthServicePortMockRecorder struct {
	mock *MockAuthServicePort
}

// NewMockAuthServicePort creates a new mock instance.
func NewMockAuthServicePort(ctrl *gomock.Controller) *MockAuthServicePort {
	mock := &MockAuthServicePort{ctrl: ctrl}
	mock.recorder = &MockAuthServicePortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServicePort) EXPECT() *MockAuthServicePortMockRecorder {
	return m.recorder
}

// LogOut mocks base method.
func (m *MockAuthServicePort) LogOut(ctx context.Context, cookie string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogOut", ctx, cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogOut indicates an expected call of LogOut.
func (mr *MockAuthServicePortMockRecorder) LogOut(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogOut", reflect.TypeOf((*MockAuthServicePort)(nil).LogOut), ctx, cookie)
}

// RefreshAccessToken mocks base method.
func (m *MockAuthServicePort) RefreshAccessToken(ctx context.Context, cookie string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshAccessToken", ctx, cookie)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshAccessToken indicates an expected call of RefreshAccessToken.
func (mr *MockAuthServicePortMockRecorder) RefreshAccessToken(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshAccessToken", reflect.TypeOf((*MockAuthServicePort)(nil).RefreshAccessToken), ctx, cookie)
}

// SignIn mocks base method.
func (m *MockAuthServicePort) SignIn(ctx context.Context, req *gen.LoginRequest) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthServicePortMockRecorder) SignIn(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthServicePort)(nil).SignIn), ctx, req)
}

// SignUpAdmin mocks base method.
func (m *MockAuthServicePort) SignUpAdmin(ctx context.Context, req *gen.CreateAdminRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpAdmin", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUpAdmin indicates an expected call of SignUpAdmin.
func (mr *MockAuthServicePortMockRecorder) SignUpAdmin(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpAdmin", reflect.TypeOf((*MockAuthServicePort)(nil).SignUpAdmin), ctx, req)
}

// SignUpCompany mocks base method.
func (m *MockAuthServicePort) SignUpCompany(ctx context.Context, req *gen.CreateCompanyRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpCompany", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUpCompany indicates an expected call of SignUpCompany.
func (mr *MockAuthServicePortMockRecorder) SignUpCompany(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpCompany", reflect.TypeOf((*MockAuthServicePort)(nil).SignUpCompany), ctx, req)
}

// SignUpStudent mocks base method.
func (m *MockAuthServicePort) SignUpStudent(ctx context.Context, req *gen.CreateStudentRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpStudent", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUpStudent indicates an expected call of SignUpStudent.
func (mr *MockAuthServicePortMockRecorder) SignUpStudent(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpStudent", reflect.TypeOf((*MockAuthServicePort)(nil).SignUpStudent), ctx, req)
}

// VerifyEmail mocks base method.
func (m *MockAuthServicePort) VerifyEmail(ctx context.Context, sid, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmail", ctx, sid, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyEmail indicates an expected call of VerifyEmail.
func (mr *MockAuthServicePortMockRecorder) VerifyEmail(ctx, sid, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmail", reflect.TypeOf((*MockAuthServicePort)(nil).VerifyEmail), ctx, sid, code)
}

// MockUserServicePort is a mock of UserServicePort interface.
type MockUserServicePort struct {
	ctrl     *gomock.Controller
	recorder *MockUserServicePortMockRecorder
}

// MockUserServicePortMockRecorder is the mock recorder for MockUserServicePort.
type MockUserServicePortMockRecorder struct {
	mock *MockUserServicePort
}

// NewMockUserServicePort creates a new mock instance.
func NewMockUserServicePort(ctrl *gomock.Controller) *MockUserServicePort {
	mock := &MockUserServicePort{ctrl: ctrl}
	mock.recorder = &MockUserServicePortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServicePort) EXPECT() *MockUserServicePortMockRecorder {
	return m.recorder
}

// DeleteCompany mocks base method.
func (m *MockUserServicePort) DeleteCompany(ctx context.Context, userId, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", ctx, userId, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockUserServicePortMockRecorder) DeleteCompany(ctx, userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockUserServicePort)(nil).DeleteCompany), ctx, userId, id)
}

// DeleteStudent mocks base method.
func (m *MockUserServicePort) DeleteStudent(ctx context.Context, userId, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", ctx, userId, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockUserServicePortMockRecorder) DeleteStudent(ctx, userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockUserServicePort)(nil).DeleteStudent), ctx, userId, id)
}

// GetAllCompany mocks base method.
func (m *MockUserServicePort) GetAllCompany(ctx context.Context, userId int64) ([]*gen.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCompany", ctx, userId)
	ret0, _ := ret[0].([]*gen.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCompany indicates an expected call of GetAllCompany.
func (mr *MockUserServicePortMockRecorder) GetAllCompany(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCompany", reflect.TypeOf((*MockUserServicePort)(nil).GetAllCompany), ctx, userId)
}

// GetApprovedCompany mocks base method.
func (m *MockUserServicePort) GetApprovedCompany(ctx context.Context, userId int64, search string) ([]*gen.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApprovedCompany", ctx, userId, search)
	ret0, _ := ret[0].([]*gen.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApprovedCompany indicates an expected call of GetApprovedCompany.
func (mr *MockUserServicePortMockRecorder) GetApprovedCompany(ctx, userId, search interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApprovedCompany", reflect.TypeOf((*MockUserServicePort)(nil).GetApprovedCompany), ctx, userId, search)
}

// GetCompanyByID mocks base method.
func (m *MockUserServicePort) GetCompanyByID(ctx context.Context, userId, id int64) (*gen.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyByID", ctx, userId, id)
	ret0, _ := ret[0].(*gen.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyByID indicates an expected call of GetCompanyByID.
func (mr *MockUserServicePortMockRecorder) GetCompanyByID(ctx, userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyByID", reflect.TypeOf((*MockUserServicePort)(nil).GetCompanyByID), ctx, userId, id)
}

// GetCompanyMe mocks base method.
func (m *MockUserServicePort) GetCompanyMe(ctx context.Context, id int64) (*gen.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyMe", ctx, id)
	ret0, _ := ret[0].(*gen.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyMe indicates an expected call of GetCompanyMe.
func (mr *MockUserServicePortMockRecorder) GetCompanyMe(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyMe", reflect.TypeOf((*MockUserServicePort)(nil).GetCompanyMe), ctx, id)
}

// GetStudentByID mocks base method.
func (m *MockUserServicePort) GetStudentByID(ctx context.Context, userId, id int64) (*gen.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentByID", ctx, userId, id)
	ret0, _ := ret[0].(*gen.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentByID indicates an expected call of GetStudentByID.
func (mr *MockUserServicePortMockRecorder) GetStudentByID(ctx, userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentByID", reflect.TypeOf((*MockUserServicePort)(nil).GetStudentByID), ctx, userId, id)
}

// GetStudentMe mocks base method.
func (m *MockUserServicePort) GetStudentMe(ctx context.Context, id int64) (*gen.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentMe", ctx, id)
	ret0, _ := ret[0].(*gen.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentMe indicates an expected call of GetStudentMe.
func (mr *MockUserServicePortMockRecorder) GetStudentMe(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentMe", reflect.TypeOf((*MockUserServicePort)(nil).GetStudentMe), ctx, id)
}

// UpdateCompanyMe mocks base method.
func (m *MockUserServicePort) UpdateCompanyMe(ctx context.Context, id int64, req *gen.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompanyMe", ctx, id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompanyMe indicates an expected call of UpdateCompanyMe.
func (mr *MockUserServicePortMockRecorder) UpdateCompanyMe(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompanyMe", reflect.TypeOf((*MockUserServicePort)(nil).UpdateCompanyMe), ctx, id, req)
}

// UpdateCompanyStatus mocks base method.
func (m *MockUserServicePort) UpdateCompanyStatus(ctx context.Context, userId, id int64, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompanyStatus", ctx, userId, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompanyStatus indicates an expected call of UpdateCompanyStatus.
func (mr *MockUserServicePortMockRecorder) UpdateCompanyStatus(ctx, userId, id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompanyStatus", reflect.TypeOf((*MockUserServicePort)(nil).UpdateCompanyStatus), ctx, userId, id, status)
}

// UpdateStudentMe mocks base method.
func (m *MockUserServicePort) UpdateStudentMe(ctx context.Context, id int64, req *gen.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudentMe", ctx, id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStudentMe indicates an expected call of UpdateStudentMe.
func (mr *MockUserServicePortMockRecorder) UpdateStudentMe(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudentMe", reflect.TypeOf((*MockUserServicePort)(nil).UpdateStudentMe), ctx, id, req)
}
