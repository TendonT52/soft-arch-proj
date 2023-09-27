package test

import (
	"context"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/server"
	mock "github.com/TikhampornSky/go-auth-verifiedMail/test/mock_port"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateStudentSuccess(t *testing.T) {
	u := pbv1.CreateStudentRequest{
		Name:            "mock-student-name",
		Email:           "6331122233@student.chula.ac.th",
		Password:        "mypassword",
		PasswordConfirm: "mypassword",
		Description:     "mock-student-description",
		Faculty:         "mock-student-faculty",
		Major:           "mock-student-major",
		Year:            4,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpStudent(gomock.Any(), &u).Return(int64(1), nil)

	s := server.NewAuthServer(m)
	r, err := s.CreateStudent(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(201), r.Status)
}

func TestCreateStudentPasswordNotMatch(t *testing.T) {
	u := pbv1.CreateStudentRequest{
		Name:            "mock-student-name-1",
		Email:           "6331122234@student.chula.ac.th",
		Password:        "mypassword-1",
		PasswordConfirm: "mypassword-",
		Description:     "mock-student-description-1",
		Faculty:         "mock-student-faculty-1",
		Major:           "mock-student-major-1",
		Year:            4,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpStudent(gomock.Any(), &u).Return(int64(0), domain.ErrPasswordNotMatch)

	s := server.NewAuthServer(m)
	r, err := s.CreateStudent(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestCreateStudentNotChulaEmail(t *testing.T) {
	u := pbv1.CreateStudentRequest{
		Name:            "mock-student-name-2",
		Email:           "weird-email",
		Password:        "mypassword-2",
		PasswordConfirm: "mypassword-2",
		Description:     "mock-student-description-2",
		Faculty:         "mock-student-faculty-2",
		Major:           "mock-student-major-2",
		Year:            4,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpStudent(gomock.Any(), &u).Return(int64(0), domain.ErrNotChulaStudentEmail)

	s := server.NewAuthServer(m)
	r, err := s.CreateStudent(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestCreateStudentDuplicateEmail(t *testing.T) {
	u := pbv1.CreateStudentRequest{
		Name:            "mock-student-name-3",
		Email:           "duplicate-email",
		Password:        "mypassword-3",
		PasswordConfirm: "mypassword-3",
		Description:     "mock-student-description-3",
		Faculty:         "mock-student-faculty-3",
		Major:           "mock-student-major-3",
		Year:            4,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpStudent(gomock.Any(), &u).Return(int64(0), domain.ErrDuplicateEmail)

	s := server.NewAuthServer(m)
	r, err := s.CreateStudent(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestCreateCompanySuccess(t *testing.T) {
	u := pbv1.CreateCompanyRequest{
		Name:            "mock-company-name",
		Email:           "mock-company@email.com",
		Password:        "mypassword",
		PasswordConfirm: "mypassword",
		Description:     "mock-company-description",
		Location:        "mock-company-location",
		Phone:           "mock-company-phone",
		Category:        "mock-company-category",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpCompany(gomock.Any(), &u).Return(int64(1), nil)

	s := server.NewAuthServer(m)
	r, err := s.CreateCompany(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(201), r.Status)
}

func TestCreateCompanyPasswordNotMatch(t *testing.T) {
	u := pbv1.CreateCompanyRequest{
		Name:            "mock-company-name-1",
		Email:           "mock-company-email-1",
		Password:        "mypassword-1",
		PasswordConfirm: "mypassword-",
		Description:     "mock-company-description-1",
		Location:        "mock-company-location-1",
		Phone:           "mock-company-phone-1",
		Category:        "mock-company-category-1",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpCompany(gomock.Any(), &u).Return(int64(0), domain.ErrPasswordNotMatch)

	s := server.NewAuthServer(m)
	r, err := s.CreateCompany(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestCreateCompanyDuplicateEmail(t *testing.T) {
	u := pbv1.CreateCompanyRequest{
		Name:            "mock-company-name-2",
		Email:           "duplicate-mock-company-email-2",
		Password:        "mypassword-2",
		PasswordConfirm: "mypassword-2",
		Description:     "mock-company-description-2",
		Location:        "mock-company-location-2",
		Phone:           "mock-company-phone-2",
		Category:        "mock-company-category-2",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpCompany(gomock.Any(), &u).Return(int64(0), domain.ErrDuplicateEmail)

	s := server.NewAuthServer(m)
	r, err := s.CreateCompany(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestCreateAdminSuccess(t *testing.T) {
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	u := pbv1.CreateAdminRequest{
		Email:           "mock-company@email.com",
		Password:        "mypassword",
		PasswordConfirm: "mypassword",
		AccessToken:     admin_access_token,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpAdmin(gomock.Any(), &u).Return(int64(1), nil)

	s := server.NewAuthServer(m)
	r, err := s.CreateAdmin(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(201), r.Status)
}

func TestCreateAdminPasswordNotMatch(t *testing.T) {
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	u := pbv1.CreateAdminRequest{
		Email:           "mock-admin-email-1",
		Password:        "mypassword-1",
		PasswordConfirm: "mypassword-",
		AccessToken:     admin_access_token,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpAdmin(gomock.Any(), &u).Return(int64(0), domain.ErrPasswordNotMatch)

	s := server.NewAuthServer(m)
	r, err := s.CreateAdmin(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestCreateAdminDuplicateEmail(t *testing.T) {
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	u := pbv1.CreateAdminRequest{
		Email:           "duplicate-mock-admin-email-1",
		Password:        "mypassword-1",
		PasswordConfirm: "mypassword-1",
		AccessToken:     admin_access_token,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignUpAdmin(gomock.Any(), &u).Return(int64(0), domain.ErrDuplicateEmail)

	s := server.NewAuthServer(m)
	r, err := s.CreateAdmin(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestVerifyCodeSuccess(t *testing.T) {
	mockCode := "123456"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().VerifyEmail(gomock.Any(), "1", mockCode).Return(nil)

	p := &pbv1.VerifyEmailCodeRequest{
		Code:      mockCode,
		StudentId: "1",
	}
	s := server.NewAuthServer(m)
	r, err := s.VerifyEmailCode(context.Background(), p)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestVerifyCodeAlreadyVerified(t *testing.T) {
	mockCode := "123456"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().VerifyEmail(gomock.Any(), "2", mockCode).Return(domain.ErrAlreadyVerified)

	p := &pbv1.VerifyEmailCodeRequest{
		Code:      mockCode,
		StudentId: "2",
	}
	s := server.NewAuthServer(m)
	r, err := s.VerifyEmailCode(context.Background(), p)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestLoginSuccess(t *testing.T) {
	u := pbv1.LoginRequest{
		Email:    "mock-email@email.com",
		Password: "mypassword",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignIn(gomock.Any(), &u).Return("mock_access_token", "mock_refresh_token", nil)

	s := server.NewAuthServer(m)
	r, err := s.SignIn(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestSignInWrongPassword(t *testing.T) {
	u := pbv1.LoginRequest{
		Email:    "mock-email",
		Password: "wrong-password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignIn(gomock.Any(), &u).Return("", "", domain.ErrPasswordNotMatch)

	s := server.NewAuthServer(m)
	r, err := s.SignIn(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestSignInEmailNotVerified(t *testing.T) {
	u := pbv1.LoginRequest{
		Email:    "mock-email-1",
		Password: "mock-password-1",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().SignIn(gomock.Any(), &u).Return("", "", domain.ErrNotVerified)

	s := server.NewAuthServer(m)
	r, err := s.SignIn(context.Background(), &u)
	require.NoError(t, err)
	require.Equal(t, int64(400), r.Status)
}

func TestLogOutSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().LogOut(gomock.Any(), "mock_refresh_token").Return(nil)

	p := &pbv1.LogOutRequest{
		RefreshToken: "mock_refresh_token",
	}
	s := server.NewAuthServer(m)
	r, err := s.LogOut(context.Background(), p)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestLogOutFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().LogOut(gomock.Any(), "mock_refresh_token_wrong").Return(domain.ErrInternal)

	p := &pbv1.LogOutRequest{
		RefreshToken: "mock_refresh_token_wrong",
	}
	s := server.NewAuthServer(m)
	r, err := s.LogOut(context.Background(), p)
	require.NoError(t, err)
	require.Equal(t, int64(500), r.Status)
}

func TestRefreshTokenSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().RefreshAccessToken(gomock.Any(), "mock_refresh_token").Return("mock_new_access_token", nil)

	p := &pbv1.RefreshTokenRequest{
		RefreshToken: "mock_refresh_token",
	}
	s := server.NewAuthServer(m)
	r, err := s.RefreshToken(context.Background(), p)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestRefreshTokenFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockAuthServicePort(ctrl)
	m.EXPECT().RefreshAccessToken(gomock.Any(), "mock_refresh_token_wrong").Return("", domain.ErrUserIDNotFound)

	p := &pbv1.RefreshTokenRequest{
		RefreshToken: "mock_refresh_token_wrong",
	}
	s := server.NewAuthServer(m)
	r, err := s.RefreshToken(context.Background(), p)
	require.NoError(t, err)
	require.Equal(t, int64(403), r.Status)
}
