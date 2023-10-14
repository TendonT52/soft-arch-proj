package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/JinnnDamanee/review-service/domain"
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

func (r *reviewRepository) CreateReview(ctx context.Context, userID int64, review *pbv1.CreatedReview) (int64, error) {
	current_timestamp := time.Now().Unix()

	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO reviews (uid, cid, title, description, rating, anonymous, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING rid")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRowContext(ctx, userID, review.Cid, review.Title, review.Description, review.Rating, review.IsAnonymous, current_timestamp, current_timestamp).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *reviewRepository) GetReviewByID(ctx context.Context, reviewID int64) (*pbv1.Review, error) {
	query := `SELECT uid, cid, title, description, rating, anonymous, updated_at FROM reviews WHERE rid = $1`

	var uid int64
	var cid int64
	var IsAnonymous bool
	var review pbv1.Review
	err := r.db.QueryRowContext(ctx, query, reviewID).Scan(&uid, &cid, &review.Title, &review.Description, &review.Rating, &IsAnonymous, &review.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrReviewNotFound
	}
	if err != nil {
		return nil, err
	}
	review.Owner = &pbv1.Owner{}
	review.Company = &pbv1.ReviewdCompany{}

	if IsAnonymous {
		review.Owner.Id = 0
	} else {
		review.Owner.Id = uid
	}

	review.Company.Id = cid
	review.Id = reviewID
	return &review, nil
}

func (r *reviewRepository) GetReviewsByCompany(ctx context.Context, companyID int64) ([]*pbv1.ReviewCompany, error) {
	query := `SELECT rid, uid, title, description, rating, anonymous, updated_at FROM reviews WHERE cid = $1`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*pbv1.ReviewCompany
	for rows.Next() {
		var uid int64
		var IsAnonymous bool
		var review pbv1.ReviewCompany
		err := rows.Scan(&review.Id, &uid, &review.Title, &review.Description, &review.Rating, &IsAnonymous, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if IsAnonymous {
			review.Owner.Id = 0
		} else {
			review.Owner.Id = uid
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewOwner(ctx context.Context, reviewID int64) (int64, error) {
	query := `SELECT uid FROM reviews WHERE rid = $1`
	row := r.db.QueryRowContext(ctx, query, reviewID)

	var uid int64
	err := row.Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (r *reviewRepository) UpdateReview(ctx context.Context, review *pbv1.UpdatedReview, rid int64) error {
	query := `UPDATE reviews SET title = $1, description = $2, rating = $3, anonymous = $4, updated_at = $5 WHERE rid = $6`
	_, err := r.db.ExecContext(ctx, query, review.Title, review.Description, review.Rating, review.IsAnonymous, time.Now().Unix(), rid)
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) GetReviewsByUser(ctx context.Context, userID int64) ([]*pbv1.MyReview, error) {
	panic("NEED Implement from Jindamanee")
	// Similar to GetReviewsByCompany function, but doesn't need to check anonymous
}

func (r *reviewRepository) DeleteReview(ctx context.Context, reviewID int64) error {
	panic("NEED Implement from Jindamanee")
}
