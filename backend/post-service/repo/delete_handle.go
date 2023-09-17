package repo

import (
	"context"
	"database/sql"
)

func (r *postRepository) deleteOpenPositions(ctx context.Context, tx *sql.Tx, openPositions *[]string, postId int64) error {
	// Check if there is any open position in posts_open_positions table
	query0 := `SELECT oid FROM open_positions WHERE title = $1`
	stmt0, err := tx.PrepareContext(ctx, query0)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt0.Close()

	query1 := `SELECT pid FROM posts_open_positions WHERE oid = $1 AND pid != $2 LIMIT 1`
	stmt1, err := tx.PrepareContext(ctx, query1)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt1.Close()

	// Delete from open_positions
	query2 := `DELETE FROM open_positions WHERE title = $1`
	stmt2, err := tx.PrepareContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt2.Close()

	// Delete from posts_open_positions
	query3 := `DELETE FROM posts_open_positions WHERE pid = $1 AND oid = $2`
	stmt3, err := tx.PrepareContext(ctx, query3)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt3.Close()

	for _, title := range *openPositions {
		var openPositionId int64
		err = stmt0.QueryRowContext(ctx, title).Scan(&openPositionId)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmt3.ExecContext(ctx, postId, openPositionId)
		if err != nil {
			tx.Rollback()
			return err
		}

		var tmpId int64 = 0
		err = stmt1.QueryRowContext(ctx, openPositionId, postId).Scan(&tmpId)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}
		if err == sql.ErrNoRows || tmpId == 0 {
			_, err = stmt2.ExecContext(ctx, title)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

func (r *postRepository) deleteRequiredSkills(ctx context.Context, tx *sql.Tx, requiredSkills *[]string, postId int64) error {
	// Check if there is any required skills in posts_required_skills table
	query0 := `SELECT sid FROM required_skills WHERE title = $1`
	stmt0, err := tx.PrepareContext(ctx, query0)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt0.Close()

	query1 := `SELECT pid FROM posts_required_skills WHERE sid = $1 AND pid != $2 LIMIT 1`
	stmt1, err := tx.PrepareContext(ctx, query1)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt1.Close()

	// Delete from required_skills
	query2 := `DELETE FROM required_skills WHERE title = $1`
	stmt2, err := tx.PrepareContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt2.Close()

	// Delete from posts_required_skills
	query3 := `DELETE FROM posts_required_skills WHERE pid = $1 AND sid = $2`
	stmt3, err := tx.PrepareContext(ctx, query3)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt3.Close()

	for _, title := range *requiredSkills {
		var requiredSkillId int64
		err = stmt0.QueryRowContext(ctx, title).Scan(&requiredSkillId)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmt3.ExecContext(ctx, postId, requiredSkillId)
		if err != nil {
			tx.Rollback()
			return err
		}

		var tmpId int64 = 0
		err = stmt1.QueryRowContext(ctx, requiredSkillId, postId).Scan(&tmpId)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}
		if err == sql.ErrNoRows || tmpId == 0 {
			_, err = stmt2.ExecContext(ctx, title)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

func (r *postRepository) deleteBenefits(ctx context.Context, tx *sql.Tx, benefits *[]string, postId int64) error {
	// Check if there is any benefits in posts_benefits table
	query0 := `SELECT bid FROM benefits WHERE title = $1`
	stmt0, err := tx.PrepareContext(ctx, query0)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt0.Close()

	query1 := `SELECT pid FROM posts_benefits WHERE bid = $1 AND pid != $2 LIMIT 1`
	stmt1, err := tx.PrepareContext(ctx, query1)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt1.Close()

	// Delete from benefits
	query2 := `DELETE FROM benefits WHERE title = $1`
	stmt2, err := tx.PrepareContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt2.Close()

	// Delete from posts_benefits
	query3 := `DELETE FROM posts_benefits WHERE pid = $1 AND bid = $2`
	stmt3, err := tx.PrepareContext(ctx, query3)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt3.Close()

	for _, title := range *benefits {
		var benefitId int64
		err = stmt0.QueryRowContext(ctx, title).Scan(&benefitId)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmt3.ExecContext(ctx, postId, benefitId)
		if err != nil {
			tx.Rollback()
			return err
		}

		var tmpId int64 = 0
		err = stmt1.QueryRowContext(ctx, benefitId, postId).Scan(&tmpId)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}
		if tmpId == 0 || err == sql.ErrNoRows {
			_, err = stmt2.ExecContext(ctx, title)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}
