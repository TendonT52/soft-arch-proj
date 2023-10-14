package repo

import (
	"context"
	"database/sql"

	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/port"
	_ "github.com/lib/pq"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type reviewRepository struct {
	db DBTX
}

func NewReviewRepository(db DBTX) port.ReviewRepoPort {
	return &reviewRepository{db: db}
}

// CreateReview implements port.ReviewRepoPort.
func (*reviewRepository) CreateReview(ctx context.Context, review *pbv1.CreatedReview) (int64, error) {
	panic("unimplemented")
}

// DeleteReview implements port.ReviewRepoPort.
func (*reviewRepository) DeleteReview(ctx context.Context, reviewID int64) error {
	panic("unimplemented")
}

// GetReviewByID implements port.ReviewRepoPort.
func (*reviewRepository) GetReviewByID(ctx context.Context, reviewID int64) (*pbv1.Review, error) {
	panic("unimplemented")
}

// GetReviewsByCompany implements port.ReviewRepoPort.
func (*reviewRepository) GetReviewsByCompany(ctx context.Context, companyID int64) ([]*pbv1.Review, error) {
	panic("unimplemented")
}

// GetReviewsByUser implements port.ReviewRepoPort.
func (*reviewRepository) GetReviewsByUser(ctx context.Context, userID int64) ([]*pbv1.Review, error) {
	panic("unimplemented")
}

// UpdateReview implements port.ReviewRepoPort.
func (*reviewRepository) UpdateReview(ctx context.Context, review *pbv1.UpdatedReview) (*pbv1.Review, error) {
	panic("unimplemented")
}
