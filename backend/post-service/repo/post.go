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
	err = r.insertIntoOpenPositions(ctx, tx, post.OpenPositions, postId)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into required_skills and posts_required_skills
	err = r.insertIntoRequiredSkills(ctx, tx, post.RequiredSkills, postId)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into benefits and posts_benefits
	err = r.insertIntoBenefits(ctx, tx, post.Benefits, postId)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	return postId, nil
}

func (r *postRepository) GetPost(ctx context.Context, postId int64) (*pbv1.Post, error) { // Internal use only
	query := "SELECT topic, description, period, how_to FROM posts WHERE pid = $1"
	var topic, description, period, howTo string
	err := r.db.QueryRowContext(ctx, query, postId).Scan(&topic, &description, &period, &howTo)
	if err == sql.ErrNoRows {
		return nil, domain.ErrPostNotFound
	}
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}

	// Find open_positions
	query1 := "SELECT title FROM open_positions WHERE oid IN (SELECT oid FROM posts_open_positions WHERE pid = $1)"
	rows, err := r.db.QueryContext(ctx, query1, postId)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	var openPositions []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		openPositions = append(openPositions, title)
	}

	// Find required_skills
	query2 := "SELECT title FROM required_skills WHERE sid IN (SELECT sid FROM posts_required_skills WHERE pid = $1)"
	rows, err = r.db.QueryContext(ctx, query2, postId)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	var requiredSkills []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		requiredSkills = append(requiredSkills, title)
	}

	// Find benefits
	query3 := "SELECT title FROM benefits WHERE bid IN (SELECT bid FROM posts_benefits WHERE pid = $1)"
	rows, err = r.db.QueryContext(ctx, query3, postId)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	var benefits []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		benefits = append(benefits, title)
	}

	post := &pbv1.Post{
		Topic:          topic,
		Description:    description,
		Period:         period,
		HowTo:          howTo,
		OpenPositions:  openPositions,
		RequiredSkills: requiredSkills,
		Benefits:       benefits,
	}

	return post, nil
}

func (r *postRepository) GetPosts(ctx context.Context, search string) ([]*pbv1.Post, error) { // Search !!!!
	return nil, nil
}

func (r *postRepository) GetOwner(ctx context.Context, postId int64) (int64, error) {
	query := "SELECT uid FROM posts WHERE pid = $1"
	var userId int64
	err := r.db.QueryRowContext(ctx, query, postId).Scan(&userId)
	if err == sql.ErrNoRows {
		return 0, domain.ErrUserIdNotFound
	}
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}
	return userId, nil
}

func (r *postRepository) UpdatePost(ctx context.Context, postId int64, post *domain.Post) error {
	current_timestamp := time.Now().Unix()

	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Update Table posts
	query := "UPDATE posts SET topic = $1, description = $2, period = $3, how_to = $4, updated_at = $5 WHERE pid = $6"
	_, err = tx.ExecContext(ctx, query, post.Topic, post.Description, post.Period, post.HowTo, current_timestamp, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// ================================================================
	// Find and Delete from table posts_open_positions
	query1 := "DELETE FROM posts_open_positions WHERE pid = $1 RETURNING oid"
	var openPositionIds []int64
	rows, err := tx.QueryContext(ctx, query1, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	for rows.Next() {
		var openPositionId int64
		err = rows.Scan(&openPositionId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		openPositionIds = append(openPositionIds, openPositionId)
	}

	// Find and Delete from table posts_required_skills
	query2 := "DELETE FROM posts_required_skills WHERE pid = $1 RETURNING sid"
	var requiredSkillIds []int64
	rows, err = tx.QueryContext(ctx, query2, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	for rows.Next() {
		var requiredSkillId int64
		err = rows.Scan(&requiredSkillId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		requiredSkillIds = append(requiredSkillIds, requiredSkillId)
	}

	// Find and Delete from table posts_benefits
	query3 := "DELETE FROM posts_benefits WHERE pid = $1 RETURNING bid"
	var benefitIds []int64
	rows, err = tx.QueryContext(ctx, query3, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	for rows.Next() {
		var benefitId int64
		err = rows.Scan(&benefitId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		benefitIds = append(benefitIds, benefitId)
	}
	// ================================================================

	// ================================================================
	// Delete from table open_positions
	query4 := "DELETE FROM open_positions WHERE oid = $1"
	stmt1, err := tx.PrepareContext(ctx, query4)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt1.Close()
	for _, openPositionId := range openPositionIds {
		_, err = stmt1.ExecContext(ctx, openPositionId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	// Delete from table required_skills
	query5 := "DELETE FROM required_skills WHERE sid = $1"
	stmt2, err := tx.PrepareContext(ctx, query5)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt2.Close()
	for _, requiredSkillId := range requiredSkillIds {
		_, err = stmt2.ExecContext(ctx, requiredSkillId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	// Delete from table benefits
	query6 := "DELETE FROM benefits WHERE bid = $1"
	stmt3, err := tx.PrepareContext(ctx, query6)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt3.Close()
	for _, benefitId := range benefitIds {
		_, err = stmt3.ExecContext(ctx, benefitId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}
	// ================================================================

	// Insert into open_positions and posts_open_positions
	err = r.insertIntoOpenPositions(ctx, tx, post.OpenPositions, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into required_skills and posts_required_skills
	err = r.insertIntoRequiredSkills(ctx, tx, post.RequiredSkills, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into benefits and posts_benefits
	err = r.insertIntoBenefits(ctx, tx, post.Benefits, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *postRepository) DeletePost(ctx context.Context, postId int64) error {
	return nil
}
