package port

import (
	"context"

	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
)

type UserRepoPort interface {
	CreateStudent(ctx context.Context, user *pbv1.CreateStudentRequest, code string) error
	CreateCompany(ctx context.Context, user *pbv1.CreateCompanyRequest) error
	CreateAdmin(ctx context.Context, user *pbv1.CreateAdminRequest) error

	UpdateVerificationCode(ctx context.Context, verification_code string) error
	GetPassword(ctx context.Context, req *pbv1.LoginRequest) (int64, string, error)
	CheckUserIDExist(ctx context.Context, id int64) (string, error)
	CheckEmailExist(ctx context.Context, email string) error
	CheckIfAdmin(ctx context.Context, id int64) error

	GetStudentByID(ctx context.Context, id int64) (*pbv1.Student, error)
	GetCompanyByID(ctx context.Context, id int64) (*pbv1.Company, error)
	GetAllCompany(ctx context.Context) ([]*pbv1.Company, error)
	GetApprovedCompany(ctx context.Context, search string) ([]*pbv1.Company, error)

	UpdateStudentByID(ctx context.Context, id int64, req *pbv1.Student) error
	UpdateCompanyByID(ctx context.Context, id int64, req *pbv1.Company) error

	UpdateCompanyStatus(ctx context.Context, id int64, status string) error

	DeleteStudent(ctx context.Context, id int64) error
	DeleteCompany(ctx context.Context, id int64) error

	SetValueRedis(ctx context.Context, key string, value string) error
	GetValueRedis(ctx context.Context, key string) (string, error)
}
