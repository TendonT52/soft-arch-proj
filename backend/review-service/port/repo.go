package port

import (
	"context"
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

type ReviewRepoPort interface {
	CreateReview(ctx context.Context, review *pbv1.CreatedReview) (int64, error)
	GetReviewsByCompany(ctx context.Context, companyID int64) ([]*pbv1.Review, error)
	GetReviewsByUser(ctx context.Context, userID int64) ([]*pbv1.Review, error)
	GetReviewByID(ctx context.Context, reviewID int64) (*pbv1.Review, error)
	UpdateReview(ctx context.Context, review *pbv1.UpdatedReview) (*pbv1.Review, error)
	DeleteReview(ctx context.Context, reviewID int64) error
}
