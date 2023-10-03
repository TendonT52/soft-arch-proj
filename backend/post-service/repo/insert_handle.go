package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/TikhampornSky/go-post-service/domain"
)

func (r *postRepository) insertIntoOpenPositions(ctx context.Context, tx *sql.Tx, openPositions *[]string, postId int64) error {
	current_timestamp := time.Now().Unix()
	// Check if open_positions table has the same title
	query := "SELECT oid FROM open_positions WHERE title = $1"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
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

	for _, title := range *openPositions {
		var openPositionId int64
		err = stmt.QueryRowContext(ctx, title).Scan(&openPositionId)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		if err == sql.ErrNoRows {
			err = stmt1.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&openPositionId)
			if err != nil {
				tx.Rollback()
				return domain.ErrInternal.From(err.Error(), err)
			}
		}

		_, err = stmt2.ExecContext(ctx, postId, openPositionId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	return nil
}

func (r *postRepository) insertIntoRequiredSkills(ctx context.Context, tx *sql.Tx, requiredSkills *[]string, postId int64) error {
	current_timestamp := time.Now().Unix()
	// Check if required_skills table has the same title
	query := "SELECT sid FROM required_skills WHERE title = $1"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	// Insert into required_skills
	query1 := "INSERT INTO required_skills (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING sid"
	stmt1, err := tx.PrepareContext(ctx, query1)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt1.Close()

	// Insert into posts_required_skills
	query2 := "INSERT INTO posts_required_skills (pid, sid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt2, err := tx.PrepareContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt2.Close()

	for _, title := range *requiredSkills {
		var requiredSkillId int64
		err = stmt.QueryRowContext(ctx, title).Scan(&requiredSkillId)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		if err == sql.ErrNoRows {
			err = stmt1.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&requiredSkillId)
			if err != nil {
				tx.Rollback()
				return domain.ErrInternal.From(err.Error(), err)
			}
		}

		_, err = stmt2.ExecContext(ctx, postId, requiredSkillId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	return nil
}

func (r *postRepository) insertIntoBenefits(ctx context.Context, tx *sql.Tx, benefits *[]string, postId int64) error {
	current_timestamp := time.Now().Unix()
	// Check if benefits table has the same title
	query := "SELECT bid FROM benefits WHERE title = $1"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into benefits
	query1 := "INSERT INTO benefits (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING bid"
	stmt1, err := tx.PrepareContext(ctx, query1)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt1.Close()

	// Insert into posts_benefits
	query2 := "INSERT INTO posts_benefits (pid, bid, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt2, err := tx.PrepareContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt2.Close()

	for _, title := range *benefits {
		var benefitId int64
		err = stmt.QueryRowContext(ctx, title).Scan(&benefitId)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		if err == sql.ErrNoRows {
			err = stmt1.QueryRowContext(ctx, title, current_timestamp, current_timestamp).Scan(&benefitId)
			if err != nil {
				tx.Rollback()
				return domain.ErrInternal.From(err.Error(), err)
			}
		}

		_, err = stmt2.ExecContext(ctx, postId, benefitId, current_timestamp, current_timestamp)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
	}

	return nil
}
