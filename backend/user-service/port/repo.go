package port

import (
	"context"

	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
)

type UserRepoPort interface {
	CreateStudent(ctx context.Context, user *pbv1.CreateStudentRequest, createTime int64) (int64, error)
	CreateCompany(ctx context.Context, user *pbv1.CreateCompanyRequest, createTime int64) (int64, error)
	CreateAdmin(ctx context.Context, user *pbv1.CreateAdminRequest, createTime int64) (int64, error)

	GetSalt(ctx context.Context, email string) (string, error)
	GetUser(ctx context.Context, req *pbv1.LoginRequest) (*pbv1.User, error)
	CheckUserIDExist(ctx context.Context, id int64) (string, error)
	CheckEmailExist(ctx context.Context, email string) error
	CheckIfAdmin(ctx context.Context, id int64) error

	GetStudentByID(ctx context.Context, id int64) (*pbv1.Student, error)
	GetCompanyByID(ctx context.Context, id int64) (*pbv1.Company, error)
	GetAllCompany(ctx context.Context) ([]*pbv1.Company, error)
	GetApprovedCompany(ctx context.Context, search string) ([]*pbv1.Company, error)

	UpdateStudentByID(ctx context.Context, id int64, req *pbv1.Student) error
	UpdateCompanyByID(ctx context.Context, id int64, req *pbv1.Company) error

	UpdateStudentStatus(ctx context.Context, email string, verified bool) error
	UpdateCompanyStatus(ctx context.Context, id int64, status string) error

	DeleteStudent(ctx context.Context, id int64) error
	DeleteCompany(ctx context.Context, id int64) error
	DeleteCompanies(ctx context.Context) error

	SetValueRedis(ctx context.Context, key string, value string) error
	GetValueRedis(ctx context.Context, key string) (string, error)
}
