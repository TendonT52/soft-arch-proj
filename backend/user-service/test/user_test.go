package test

import (
	"context"
	"testing"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/server"
	mock "github.com/TikhampornSky/go-auth-verifiedMail/test/mock_port"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func createMockToken(t *testing.T, id int64, role string) string {
	config, _ := config.LoadConfig("..")
	mock_token, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: id,
		Role:   role,
	})
	require.NoError(t, err)
	return mock_token
}

// Student Zone
func TestGetStudentMeSuccess(t *testing.T) {
	req := &pbv1.GetStudentMeRequest{
		AccessToken: createMockToken(t, 1, domain.StudentRole),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetStudentMe(gomock.Any(), int64(1)).Return(&pbv1.Student{
		Id:          1,
		Name:        "mock-student-name",
		Email:       "mock-student-email",
		Description: "mock-student-description",
		Faculty:     "mock-student-faculty",
		Major:       "mock-student-major",
		Year:        4,
	}, nil)

	s := server.NewUserServer(m)
	r, err := s.GetStudentMe(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestGetStudentSuccess(t *testing.T) {
	req := &pbv1.GetStudentRequest{
		AccessToken: createMockToken(t, 2, domain.StudentRole),
		Id:          222,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetStudentByID(gomock.Any(), int64(2), int64(222)).Return(&pbv1.Student{
		Id:          222,
		Name:        "mock-student-name-2",
		Email:       "mock-student-email-2",
		Description: "mock-student-description-2",
		Faculty:     "mock-student-faculty-2",
		Major:       "mock-student-major-2",
		Year:        4,
	}, nil)

	s := server.NewUserServer(m)
	r, err := s.GetStudent(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestGetStudentIDNotFound(t *testing.T) {
	req := &pbv1.GetStudentRequest{
		AccessToken: createMockToken(t, 22, domain.StudentRole),
		Id:          2222,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetStudentByID(gomock.Any(), int64(22), int64(2222)).Return(&pbv1.Student{}, domain.ErrUserIDNotFound)

	s := server.NewUserServer(m)
	r, err := s.GetStudent(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(404), r.Status)
}

func TestUpdateStudentSuccess(t *testing.T) {
	req := &pbv1.UpdateStudentRequest{
		AccessToken: createMockToken(t, 3, domain.StudentRole),
		Student: &pbv1.Student{
			Id:          3,
			Name:        "mock-student-name-3",
			Email:       "mock-student-email-3",
			Description: "mock-student-description-3",
			Faculty:     "mock-student-faculty-3",
			Major:       "mock-student-major-3",
			Year:        4,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().UpdateStudentMe(gomock.Any(), int64(3), req.Student).Return(nil)

	s := server.NewUserServer(m)
	r, err := s.UpdateStudent(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestUpdateStudentUnAuthorized(t *testing.T) {
	req := &pbv1.UpdateStudentRequest{
		AccessToken: createMockToken(t, 33, domain.StudentRole),
		Student: &pbv1.Student{
			Id:          33,
			Name:        "mock-student-name-33",
			Email:       "mock-student-email-33",
			Description: "mock-student-description-33",
			Faculty:     "mock-student-faculty-33",
			Major:       "mock-student-major-33",
			Year:        4,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().UpdateStudentMe(gomock.Any(), int64(33), req.Student).Return(domain.ErrNotAuthorized)

	s := server.NewUserServer(m)
	r, err := s.UpdateStudent(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(401), r.Status)
}

// Company Zone
func TestGetCompanyMeSuccess(t *testing.T) {
	req := &pbv1.GetCompanyMeRequest{
		AccessToken: createMockToken(t, 4, domain.CompanyRole),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetCompanyMe(gomock.Any(), int64(4)).Return(&pbv1.Company{
		Id:          4,
		Name:        "mock-company-name",
		Email:       "mock-company-email",
		Description: "mock-company-description",
		Location:    "mock-company-location",
		Phone:       "mock-company-phone",
		Category:    "mock-company-category",
	}, nil)

	s := server.NewUserServer(m)
	r, err := s.GetCompanyMe(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestGetCompanySuccess(t *testing.T) {
	req := &pbv1.GetCompanyRequest{
		AccessToken: createMockToken(t, 5, domain.CompanyRole),
		Id:          555,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetCompanyByID(gomock.Any(), int64(5), int64(555)).Return(&pbv1.Company{
		Id:          555,
		Name:        "mock-company-name-2",
		Email:       "mock-company-email-2",
		Description: "mock-company-description-2",
		Location:    "mock-company-location-2",
		Phone:       "mock-company-phone-2",
		Category:    "mock-company-category-2",
	}, nil)

	s := server.NewUserServer(m)
	r, err := s.GetCompany(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestGetCompanyIDNotFound(t *testing.T) {
	req := &pbv1.GetCompanyRequest{
		AccessToken: createMockToken(t, 55, domain.CompanyRole),
		Id:          5555,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetCompanyByID(gomock.Any(), int64(55), int64(5555)).Return(&pbv1.Company{}, domain.ErrUserIDNotFound)

	s := server.NewUserServer(m)
	r, err := s.GetCompany(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(404), r.Status)
}

func TestUpdateCompanySuccess(t *testing.T) {
	req := &pbv1.UpdateCompanyRequest{
		AccessToken: createMockToken(t, 6, domain.CompanyRole),
		Company: &pbv1.Company{
			Id:          6,
			Name:        "mock-company-name-3",
			Email:       "mock-company-email-3",
			Description: "mock-company-description-3",
			Location:    "mock-company-location-3",
			Phone:       "mock-company-phone-3",
			Category:    "mock-company-category-3",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().UpdateCompanyMe(gomock.Any(), int64(6), req.Company).Return(nil)

	s := server.NewUserServer(m)
	r, err := s.UpdateCompany(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestUpdateCompanyUnAuthorized(t *testing.T) {
	req := &pbv1.UpdateCompanyRequest{
		AccessToken: createMockToken(t, 66, domain.CompanyRole),
		Company: &pbv1.Company{
			Id:          66,
			Name:        "mock-company-name-33",
			Email:       "mock-company-email-33",
			Description: "mock-company-description-33",
			Location:    "mock-company-location-33",
			Phone:       "mock-company-phone-33",
			Category:    "mock-company-category-33",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().UpdateCompanyMe(gomock.Any(), int64(66), req.Company).Return(domain.ErrNotAuthorized)

	s := server.NewUserServer(m)
	r, err := s.UpdateCompany(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(401), r.Status)
}

func TestListCompaniesSuccess(t *testing.T) {
	req := &pbv1.ListCompaniesRequest{
		AccessToken: createMockToken(t, 7, domain.AdminRole),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetAllCompany(gomock.Any(), int64(7)).Return([]*pbv1.Company{
		{
			Id:          77,
			Name:        "mock-company-name-7",
			Email:       "mock-company-email-7",
			Description: "mock-company-description-7",
			Location:    "mock-company-location-7",
			Phone:       "mock-company-phone-7",
			Category:    "mock-company-category-7",
		},
		{
			Id:          777,
			Name:        "mock-company-name-77",
			Email:       "mock-company-email-77",
			Description: "mock-company-description-77",
			Location:    "mock-company-location-77",
			Phone:       "mock-company-phone-77",
			Category:    "mock-company-category-77",
		},
	}, nil)

	s := server.NewUserServer(m)
	r, err := s.ListCompanies(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestListCompaniesNotAuthorized(t *testing.T) {
	req := &pbv1.ListCompaniesRequest{
		AccessToken: createMockToken(t, 77, domain.AdminRole),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetAllCompany(gomock.Any(), int64(77)).Return(nil, domain.ErrNotAuthorized)

	s := server.NewUserServer(m)
	r, err := s.ListCompanies(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(401), r.Status)
}

func TestListApprovedCompaniesSuccess(t *testing.T) {
	req := &pbv1.ListApprovedCompaniesRequest{
		AccessToken: createMockToken(t, 8, domain.AdminRole),
		Search:      "search",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().GetApprovedCompany(gomock.Any(), int64(8), "search").Return([]*pbv1.Company{
		{
			Id:          88,
			Name:        "mock-company-name-8",
			Email:       "mock-company-email-8",
			Description: "mock-company-description-8",
			Location:    "mock-company-location-8",
			Phone:       "mock-company-phone-8",
			Category:    "mock-company-category-8",
		},
		{
			Id:          888,
			Name:        "mock-company-name-88",
			Email:       "mock-company-email-88",
			Description: "mock-company-description-88",
			Location:    "mock-company-location-88",
			Phone:       "mock-company-phone-88",
			Category:    "mock-company-category-88",
		},
	}, nil)

	s := server.NewUserServer(m)
	r, err := s.ListApprovedCompanies(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestApproveCompanySuccess(t *testing.T) {
	req := &pbv1.UpdateCompanyStatusRequest{
		AccessToken: createMockToken(t, 9, domain.AdminRole),
		Id:          99,
		Status:      "Approve",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().UpdateCompanyStatus(gomock.Any(), int64(9), int64(99), "Approve").Return(nil)

	s := server.NewUserServer(m)
	r, err := s.UpdateCompanyStatus(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestRejectCompanySuccess(t *testing.T) {
	req := &pbv1.UpdateCompanyStatusRequest{
		AccessToken: createMockToken(t, 9, domain.AdminRole),
		Id:          999,
		Status:      "Reject",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().UpdateCompanyStatus(gomock.Any(), int64(9), int64(999), "Reject").Return(nil)

	s := server.NewUserServer(m)
	r, err := s.UpdateCompanyStatus(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(200), r.Status)
}

func TestUpdateCompanyStatusUnAuthorized(t *testing.T) {
	req := &pbv1.UpdateCompanyStatusRequest{
		AccessToken: createMockToken(t, 99, domain.StudentRole),
		Id:          9999,
		Status:      "Approve",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUserServicePort(ctrl)
	m.EXPECT().UpdateCompanyStatus(gomock.Any(), int64(99), int64(9999), "Approve").Return(domain.ErrNotAuthorized)

	s := server.NewUserServer(m)
	r, err := s.UpdateCompanyStatus(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, int64(401), r.Status)
}
