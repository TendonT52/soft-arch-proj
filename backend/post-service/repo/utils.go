package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/TikhampornSky/go-post-service/domain"
)

func (r *postRepository) insertIntoOpenPositions(ctx context.Context, tx *sql.Tx, openPositions []string, postId int64) error {
	current_timestamp := time.Now().Unix()
	// Insert into open_positions
	query1 := "INSERT INTO open_positions (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING oid"
	stmt1, err := tx.PrepareContext(ctx, query1)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt1.Close()

	// Insert into posts_open_positions
	query2 := "INSERT INTO posts_open_positions (pid, oid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt2, err := tx.PrepareContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt2.Close()

	for _, title := range openPositions {
		var openPositionId int64
		err = stmt1.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&openPositionId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}

		_, err = stmt2.ExecContext(ctx, postId, openPositionId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	return nil
}

func (r *postRepository) insertIntoRequiredSkills(ctx context.Context, tx *sql.Tx, requiredSkills []string, postId int64) error {
	current_timestamp := time.Now().Unix()
	// Insert into required_skills
	query3 := "INSERT INTO required_skills (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING sid"
	stmt3, err := tx.PrepareContext(ctx, query3)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt3.Close()

	// Insert into posts_required_skills
	query4 := "INSERT INTO posts_required_skills (pid, sid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt4, err := tx.PrepareContext(ctx, query4)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt4.Close()

	for _, title := range requiredSkills {
		var requiredSkillId int64
		err = stmt3.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&requiredSkillId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}

		_, err = stmt4.ExecContext(ctx, postId, requiredSkillId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	return nil
}

func (r *postRepository) insertIntoBenefits(ctx context.Context, tx *sql.Tx, benefits []string, postId int64) error {
	current_timestamp := time.Now().Unix()

	// Insert into benefits
	query5 := "INSERT INTO benefits (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING bid"
	stmt5, err := tx.PrepareContext(ctx, query5)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt5.Close()

	// Insert into posts_benefits
	query6 := "INSERT INTO posts_benefits (pid, bid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt6, err := tx.PrepareContext(ctx, query6)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt6.Close()

	for _, title := range benefits {
		var benefitId int64
		err = stmt5.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&benefitId)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}

		_, err = stmt6.ExecContext(ctx, postId, benefitId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	return nil
}
