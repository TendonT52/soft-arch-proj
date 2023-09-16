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

func (r *postRepository) findAndDeleteOpenPositions(ctx context.Context, tx *sql.Tx, postId int64) error {
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

	return nil
}

func (r *postRepository) findAndDeleteRequiredSkills(ctx context.Context, tx *sql.Tx, postId int64) error {
	// Find and Delete from table posts_required_skills
	query2 := "DELETE FROM posts_required_skills WHERE pid = $1 RETURNING sid"
	var requiredSkillIds []int64
	rows, err := tx.QueryContext(ctx, query2, postId)
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

	return nil
}

func (r *postRepository) findAndDeleteBenefits(ctx context.Context, tx *sql.Tx, postId int64) error {
	query3 := "DELETE FROM posts_benefits WHERE pid = $1 RETURNING bid"
	var benefitIds []int64
	rows, err := tx.QueryContext(ctx, query3, postId)
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
	return nil
}