package port

import (
	"context"
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

type ReviewServicePort interface {
	CreateReview(ctx context.Context, token string, review *pbv1.CreatedReview) (int64, error)
	GetReviewsByCompany(ctx context.Context, token string, companyID int64) ([]*pbv1.Review, error)
	GetReviewsByUser(ctx context.Context, token string, userID int64) ([]*pbv1.Review, error)
	GetReviewByID(ctx context.Context, token string, reviewID int64) (*pbv1.Review, error)
	UpdateReview(ctx context.Context, token string, review *pbv1.UpdatedReview) (*pbv1.Review, error)
	DeleteReview(ctx context.Context, token string, reviewID int64) error
}

type UserClientPort interface {
	GetUserProfile(ctx context.Context, req *pbv1.GetStudentRequest) (*pbv1.GetStudentRequest, error)
}