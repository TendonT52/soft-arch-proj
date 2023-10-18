package port

import (
	"context"

	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

type ReviewServicePort interface {
	CreateReview(ctx context.Context, token string, review *pbv1.CreatedReview) (int64, error)
	GetReviewsByCompany(ctx context.Context, token string, companyID int64) ([]*pbv1.ReviewCompany, error)
	GetReviewsByUser(ctx context.Context, token string, userID int64) ([]*pbv1.MyReview, error)
	GetReviewByID(ctx context.Context, token string, reviewID int64) (*pbv1.Review, error)
	UpdateReview(ctx context.Context, token string, review *pbv1.UpdatedReview, rid int64) error
	DeleteReview(ctx context.Context, token string, reviewID int64) error
}

type UserClientPort interface {
	GetStudentProfile(ctx context.Context, req *pbv1.GetStudentRequest) (*pbv1.GetStudentResponse, error)
	GetStudentProfiles(ctx context.Context, req *pbv1.GetStudentsRequest) (*pbv1.GetStudentsResponse, error)
	GetCompanyProfile(ctx context.Context, req *pbv1.GetCompanyRequest) (*pbv1.GetCompanyResponse, error)
	GetCompanyProfiles(ctx context.Context, req *pbv1.GetCompaniesRequest) (*pbv1.GetCompaniesResponse, error)
}
