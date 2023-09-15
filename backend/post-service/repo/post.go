package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
	_ "github.com/lib/pq"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type postRepository struct {
	db DBTX
}

func NewPostRepository(db DBTX) port.PostRepoPort {
	return &postRepository{db: db}
}

func (r *postRepository) CreatePost(ctx context.Context, userId int64, post *domain.Post) (int64, error) {
	current_timestamp := time.Now().Unix()

	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into Table posts
	query := "INSERT INTO posts (uid, topic, description, period, how_to, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING pid"
	var postId int64
	err = tx.QueryRowContext(ctx, query, userId, post.Topic, post.Description, post.Period, post.HowTo, current_timestamp, current_timestamp).Scan(&postId)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into open_positions and posts_open_positions
	query1 := "INSERT INTO open_positions (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING oid"
	stmt1, err := tx.PrepareContext(ctx, query1)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt1.Close()

	query2 := "INSERT INTO posts_open_positions (pid, oid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt2, err := tx.PrepareContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	for _, title := range post.OpenPositions {
		var openPositionId int64
		err = stmt1.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&openPositionId)
		if err != nil {
			tx.Rollback()
			return 0, domain.ErrInternal.From(err.Error(), err)
		}

		_, err = stmt2.ExecContext(ctx, postId, openPositionId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return 0, domain.ErrInternal.From(err.Error(), err)
		}
	}

	// Insert into required_skills and posts_required_skills
	query3 := "INSERT INTO required_skills (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING sid"
	stmt3, err := tx.PrepareContext(ctx, query3)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	query4 := "INSERT INTO posts_required_skills (pid, sid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt4, err := tx.PrepareContext(ctx, query4)

	for _, title := range post.RequiredSkills {
		var requiredSkillId int64
		err = stmt3.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&requiredSkillId)
		if err != nil {
			tx.Rollback()
			return 0, domain.ErrInternal.From(err.Error(), err)
		}

		_, err = stmt4.ExecContext(ctx, postId, requiredSkillId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return 0, domain.ErrInternal.From(err.Error(), err)
		}
	}

	// Insert into benefits and posts_benefits
	query5 := "INSERT INTO benefits (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING bid"
	stmt5, err := tx.PrepareContext(ctx, query5)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	query6 := "INSERT INTO posts_benefits (pid, bid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt6, err := tx.PrepareContext(ctx, query6)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	for _, title := range post.Benefits {
		var benefitId int64
		err = stmt5.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&benefitId)
		if err != nil {
			tx.Rollback()
			return 0, domain.ErrInternal.From(err.Error(), err)
		}

		_, err = stmt6.ExecContext(ctx, postId, benefitId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return 0, domain.ErrInternal.From(err.Error(), err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	return postId, nil
}

func (r *postRepository) GetPost(ctx context.Context, postId int64) (*pbv1.Post, error) {
	return nil, nil
}

func (r *postRepository) GetPosts(ctx context.Context, search string) ([]*pbv1.Post, error) { // Search !!!!
	return nil, nil
}

func (r *postRepository) UpdatePost(ctx context.Context, postId int64, post *pbv1.Post) error {
	return nil
}

func (r *postRepository) DeletePost(ctx context.Context, postId int64) error {
	return nil
}
