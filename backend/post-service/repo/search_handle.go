package repo

import (
	"context"
	"database/sql"

	"github.com/TikhampornSky/go-post-service/domain"
	"github.com/lib/pq"
)

func (r *postRepository) searchOpenPositions(ctx context.Context, tqueryO, searchOpenPosition string, cids *[]int64) (*[]int64, error) {
	var rows *sql.Rows
	var err error
	if searchOpenPosition == "" {
		query_open := `
			SELECT 	posts.pid
			FROM posts

			INNER JOIN posts_open_positions ON posts.pid = posts_open_positions.pid
			INNER JOIN open_positions ON posts_open_positions.oid = open_positions.oid

			WHERE
				posts.uid = ANY($1)
		`

		rows, err = r.db.QueryContext(ctx, query_open, pq.Array(*cids))
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	} else {
		query_open := `
			SELECT 	posts.pid
			FROM posts

			INNER JOIN posts_open_positions ON posts.pid = posts_open_positions.pid
			INNER JOIN open_positions ON posts_open_positions.oid = open_positions.oid

			WHERE
				(to_tsquery($1) @@ to_tsvector(open_positions.title)
				OR SIMILARITY($2, open_positions.title) > 0 )
				AND posts.uid = ANY($3)
			ORDER BY
				NULLIF(ts_rank(to_tsvector(open_positions.title), to_tsquery($1)), 0) DESC NULLS LAST,
				SIMILARITY($2, open_positions.title) DESC NULLS LAST
		`

		rows, err = r.db.QueryContext(ctx, query_open, tqueryO, searchOpenPosition, pq.Array(*cids))
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	}
	defer rows.Close()

	var order []int64
	OpenPosition := make(map[int64](*[]string))
	for rows.Next() {
		var pid int64
		err = rows.Scan(&pid)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}

		var sub []string
		s := OpenPosition[pid]
		if s == nil {
			OpenPosition[pid] = &sub
			order = append(order, pid)
		}
	}

	return &order, nil
}

func (r *postRepository) searchRequiredSkills(ctx context.Context, tqueryR, searchRequiredSkill string, cids *[]int64) (*[]int64, error) {
	var rows *sql.Rows
	var err error
	if searchRequiredSkill == "" {
		query_required := `
		SELECT 	posts.pid
		FROM posts

		INNER JOIN posts_required_skills ON posts.pid = posts_required_skills.pid
		INNER JOIN required_skills ON posts_required_skills.sid = required_skills.sid

		WHERE
			posts.uid = ANY($1)
	`

		rows, err = r.db.QueryContext(ctx, query_required, pq.Array(cids))
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	} else {
		query_required := `
		SELECT 	posts.pid
		FROM posts

		INNER JOIN posts_required_skills ON posts.pid = posts_required_skills.pid
		INNER JOIN required_skills ON posts_required_skills.sid = required_skills.sid

		WHERE
			(to_tsquery($1) @@ to_tsvector(required_skills.title)
			OR SIMILARITY($2, required_skills.title) > 0 )
			AND posts.uid = ANY($3)
		ORDER BY
			NULLIF(ts_rank(to_tsvector(required_skills.title), to_tsquery($1)), 0) DESC NULLS LAST,
			SIMILARITY($2, required_skills.title) DESC NULLS LAST
	`

		rows, err = r.db.QueryContext(ctx, query_required, tqueryR, searchRequiredSkill, pq.Array(cids))
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	}
	defer rows.Close()

	var order []int64 
	RequiredSkills := make(map[int64](*[]string))
	for rows.Next() {
		var pid int64
		err = rows.Scan(&pid)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}

		var sub []string
		s := RequiredSkills[pid]
		if s == nil {
			RequiredSkills[pid] = &sub
			order = append(order, pid)
		}
	}

	return &order, nil
}

func (r *postRepository) searchBenefits(ctx context.Context, tqueryB, searchBenefit string, cids *[]int64) (*[]int64, error) {
	var rows *sql.Rows
	var err error
	if searchBenefit == "" {
		query_benefit := `
		SELECT 	posts.pid
		FROM posts

		INNER JOIN posts_benefits ON posts.pid = posts_benefits.pid
		INNER JOIN benefits ON posts_benefits.bid = benefits.bid

		WHERE
			posts.uid = ANY($1)
	`

		rows, err = r.db.QueryContext(ctx, query_benefit, pq.Array(cids))
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	} else {
		query_benefit := `
		SELECT 	posts.pid
		FROM posts

		INNER JOIN posts_benefits ON posts.pid = posts_benefits.pid
		INNER JOIN benefits ON posts_benefits.bid = benefits.bid

		WHERE
			(to_tsquery($1) @@ to_tsvector(benefits.title)
			OR SIMILARITY($2, benefits.title) > 0 )
			AND posts.uid = ANY($3)
		ORDER BY
			NULLIF(ts_rank(to_tsvector(benefits.title), to_tsquery($1)), 0) DESC NULLS LAST,
			SIMILARITY($2, benefits.title) DESC NULLS LAST
	`

		rows, err = r.db.QueryContext(ctx, query_benefit, tqueryB, searchBenefit, pq.Array(cids))
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	}
	defer rows.Close()

	var order []int64
	Benefits := make(map[int64](*[]string))
	for rows.Next() {
		var pid int64
		err = rows.Scan(&pid)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}

		var sub []string
		s := Benefits[pid]
		if s == nil {
			Benefits[pid] = &sub
			order = append(order, pid)
		}
	}

	return &order, nil
}
