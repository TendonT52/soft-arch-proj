package service

import (
	"context"

	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/port"
)

const (
	AdminRole = "admin"
)

type reviewService struct {
	repo port.ReviewRepoPort
}

func NewReviewService(repo port.ReviewRepoPort) port.ReviewServicePort {
	return &reviewService{repo: repo}
}

// CreateReview implements port.ReviewServicePort.
func (*reviewService) CreateReview(ctx context.Context, token string, review *pbv1.CreatedReview) (int64, error) {
	panic("unimplemented")
}

// DeleteReview implements port.ReviewServicePort.
func (*reviewService) DeleteReview(ctx context.Context, token string, reviewID int64) error {
	panic("unimplemented")
}

// GetReviewByID implements port.ReviewServicePort.
func (*reviewService) GetReviewByID(ctx context.Context, token string, reviewID int64) (*pbv1.Review, error) {
	panic("unimplemented")
}

// GetReviewsByCompany implements port.ReviewServicePort.
func (*reviewService) GetReviewsByCompany(ctx context.Context, token string, companyID int64) ([]*pbv1.Review, error) {
	panic("unimplemented")
}

// GetReviewsByUser implements port.ReviewServicePort.
func (*reviewService) GetReviewsByUser(ctx context.Context, token string, userID int64) ([]*pbv1.Review, error) {
	panic("unimplemented")
}

// UpdateReview implements port.ReviewServicePort.
func (*reviewService) UpdateReview(ctx context.Context, token string, review *pbv1.UpdatedReview) (*pbv1.Review, error) {
	panic("unimplemented")
}
