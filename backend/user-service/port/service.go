package port

import (
	"context"

	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
)

type AuthServicePort interface {
	SignUpStudent(ctx context.Context, req *pbv1.CreateStudentRequest) (int64, error)
	SignUpCompany(ctx context.Context, req *pbv1.CreateCompanyRequest) (int64, error)
	SignUpAdmin(ctx context.Context, req *pbv1.CreateAdminRequest) (int64, error)
	VerifyEmail(ctx context.Context, sid, code string) error
	SignIn(ctx context.Context, req *pbv1.LoginRequest) (string, string, error)
	RefreshAccessToken(ctx context.Context, cookie string) (string, error)
	LogOut(ctx context.Context, cookie string) error
}

type UserServicePort interface {
	GetStudentMe(ctx context.Context, id int64) (*pbv1.Student, error)
	GetCompanyMe(ctx context.Context, id int64) (*pbv1.Company, error)
	GetStudentByID(ctx context.Context, userId, id int64) (*pbv1.Student, error)
	GetCompanyByID(ctx context.Context, userId, id int64) (*pbv1.Company, error)
	GetAllCompany(ctx context.Context, userId int64) ([]*pbv1.Company, error)
	GetApprovedCompany(ctx context.Context, userId int64, search string) ([]*pbv1.Company, error)
	UpdateStudentMe(ctx context.Context, id int64, req *pbv1.UpdatedStudent) error
	UpdateCompanyMe(ctx context.Context, id int64, req *pbv1.UpdatedCompany) error
	UpdateCompanyStatus(ctx context.Context, userId, id int64, status string) error
	DeleteStudent(ctx context.Context, userId, id int64) error
	DeleteCompany(ctx context.Context, userId, id int64) error
	DeleteCompanies(ctx context.Context, userId int64) error
}
