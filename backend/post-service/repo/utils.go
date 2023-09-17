package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
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

// func (r *postRepository) findAndDeleteOpenPositions(ctx context.Context, tx *sql.Tx, postId int64) error {
// 	// Find and Delete from table posts_open_positions
// 	query1 := "DELETE FROM posts_open_positions WHERE pid = $1 RETURNING oid"
// 	var openPositionIds []int64
// 	rows, err := tx.QueryContext(ctx, query1, postId)
// 	if err != nil {
// 		tx.Rollback()
// 		return domain.ErrInternal.From(err.Error(), err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var openPositionId int64
// 		err = rows.Scan(&openPositionId)
// 		if err != nil {
// 			tx.Rollback()
// 			return domain.ErrInternal.From(err.Error(), err)
// 		}
// 		openPositionIds = append(openPositionIds, openPositionId)
// 	}

// 	// Delete from table open_positions
// 	query4 := "DELETE FROM open_positions WHERE oid = $1"
// 	stmt1, err := tx.PrepareContext(ctx, query4)
// 	if err != nil {
// 		tx.Rollback()
// 		return domain.ErrInternal.From(err.Error(), err)
// 	}
// 	defer stmt1.Close()
// 	for _, openPositionId := range openPositionIds {
// 		_, err = stmt1.ExecContext(ctx, openPositionId)
// 		if err != nil {
// 			tx.Rollback()
// 			return domain.ErrInternal.From(err.Error(), err)
// 		}
// 	}

// 	return nil
// }

// func (r *postRepository) findAndDeleteRequiredSkills(ctx context.Context, tx *sql.Tx, postId int64) error {
// 	// Find and Delete from table posts_required_skills
// 	query2 := "DELETE FROM posts_required_skills WHERE pid = $1 RETURNING sid"
// 	var requiredSkillIds []int64
// 	rows, err := tx.QueryContext(ctx, query2, postId)
// 	if err != nil {
// 		tx.Rollback()
// 		return domain.ErrInternal.From(err.Error(), err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var requiredSkillId int64
// 		err = rows.Scan(&requiredSkillId)
// 		if err != nil {
// 			tx.Rollback()
// 			return domain.ErrInternal.From(err.Error(), err)
// 		}
// 		requiredSkillIds = append(requiredSkillIds, requiredSkillId)
// 	}

// 	query5 := "DELETE FROM required_skills WHERE sid = $1"
// 	stmt2, err := tx.PrepareContext(ctx, query5)
// 	if err != nil {
// 		tx.Rollback()
// 		return domain.ErrInternal.From(err.Error(), err)
// 	}
// 	defer stmt2.Close()
// 	for _, requiredSkillId := range requiredSkillIds {
// 		_, err = stmt2.ExecContext(ctx, requiredSkillId)
// 		if err != nil {
// 			tx.Rollback()
// 			return domain.ErrInternal.From(err.Error(), err)
// 		}
// 	}

// 	return nil
// }

// func (r *postRepository) findAndDeleteBenefits(ctx context.Context, tx *sql.Tx, postId int64) error {
// 	query3 := "DELETE FROM posts_benefits WHERE pid = $1 RETURNING bid"
// 	var benefitIds []int64
// 	rows, err := tx.QueryContext(ctx, query3, postId)
// 	if err != nil {
// 		tx.Rollback()
// 		return domain.ErrInternal.From(err.Error(), err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var benefitId int64
// 		err = rows.Scan(&benefitId)
// 		if err != nil {
// 			tx.Rollback()
// 			return domain.ErrInternal.From(err.Error(), err)
// 		}
// 		benefitIds = append(benefitIds, benefitId)
// 	}

// 	query6 := "DELETE FROM benefits WHERE bid = $1"
// 	stmt3, err := tx.PrepareContext(ctx, query6)
// 	if err != nil {
// 		tx.Rollback()
// 		return domain.ErrInternal.From(err.Error(), err)
// 	}
// 	defer stmt3.Close()
// 	for _, benefitId := range benefitIds {
// 		_, err = stmt3.ExecContext(ctx, benefitId)
// 		if err != nil {
// 			tx.Rollback()
// 			return domain.ErrInternal.From(err.Error(), err)
// 		}
// 	}
// 	return nil
// }

func (r *postRepository) updateOpenPosition(ctx context.Context, tx *sql.Tx, openPositions []*pbv1.Element, postId int64) error {
	var err error
	var adds []string
	var deletes []string
	for _, s := range openPositions {
		if s.IsAdd { 
			adds = append(adds, s.Value)
		} else {
			deletes = append(deletes, s.Value)
		}
	}

	err = r.insertIntoOpenPositions(ctx, tx, &adds, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	err = r.deleteOpenPositions(ctx, tx, &deletes, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *postRepository) updateRequiredSkill(ctx context.Context, tx *sql.Tx, requiredSkills []*pbv1.Element, postId int64) error {
	var err error
	var adds []string
	var deletes []string
	for _, s := range requiredSkills {
		if s.IsAdd {
			adds = append(adds, s.Value)
		} else {
			deletes = append(deletes, s.Value)
		}
	}

	err = r.insertIntoRequiredSkills(ctx, tx, &adds, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	err = r.deleteRequiredSkills(ctx, tx, &deletes, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *postRepository) updateBenefit(ctx context.Context, tx *sql.Tx, benefits []*pbv1.Element, postId int64) error {
	var err error
	var adds []string
	var deletes []string
	for _, s := range benefits {
		if s.IsAdd {
			adds = append(adds, s.Value)
		} else {
			deletes = append(deletes, s.Value)
		}
	}

	err = r.insertIntoBenefits(ctx, tx, &adds, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	err = r.deleteBenefits(ctx, tx, &deletes, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}