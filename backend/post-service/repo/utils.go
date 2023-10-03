package repo

import (
	"context"
	"database/sql"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

func (r *postRepository) updateOpenPosition(ctx context.Context, tx *sql.Tx, openPositions []*pbv1.Element, postId int64) error {
	var err error
	var adds []string
	var deletes []string

	for _, s := range openPositions {
		if s.Action == pbv1.ElementStatus_ADD {
			adds = append(adds, s.Value)
		} else if s.Action == pbv1.ElementStatus_REMOVE {
			deletes = append(deletes, s.Value)
		} else {
			// Do nothing
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
		if s.Action == pbv1.ElementStatus_ADD {
			adds = append(adds, s.Value)
		} else if s.Action == pbv1.ElementStatus_REMOVE {
			deletes = append(deletes, s.Value)
		} else {
			// Do nothing
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
		if s.Action == pbv1.ElementStatus_ADD {
			adds = append(adds, s.Value)
		} else if s.Action == pbv1.ElementStatus_REMOVE {
			deletes = append(deletes, s.Value)
		} else {
			// Do nothing
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
