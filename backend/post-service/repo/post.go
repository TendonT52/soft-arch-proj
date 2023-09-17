package repo

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
	"github.com/lib/pq"
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

func (r *postRepository) CreatePost(ctx context.Context, userId int64, post *pbv1.Post) (int64, error) {
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

func (r *postRepository) GetPost(ctx context.Context, postId int64) (*pbv1.Post, error) {
	query := "SELECT topic, description, period, how_to, uid FROM posts WHERE pid = $1"
	var topic, description, period, howTo string
	var userId int64
	err := r.db.QueryRowContext(ctx, query, postId).Scan(&topic, &description, &period, &howTo, &userId)
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
		Owner: &pbv1.PostOwner{
			Id: userId,
		},
	}

	return post, nil
}

func (r *postRepository) GetPosts(ctx context.Context, search *pbv1.SearchOptions, cids []*int64) ([]*pbv1.Post, error) {
	var searchResult *domain.IndividualSearchResult = &domain.IndividualSearchResult{
		OpenPosition:  make(map[int64](*[]domain.SubSearch)),
		RequiredSkill: make(map[int64](*[]domain.SubSearch)),
		Benefits:      make(map[int64](*[]domain.SubSearch)),
	}

	// Create
	parts := strings.Fields(search.SearchBenefit)
	tqueryB := strings.Join(parts, " | ")

	parts = strings.Fields(search.SearchOpenPosition)
	tqueryO := strings.Join(parts, " | ")

	parts = strings.Fields(search.SearchRequiredSkill)
	tqueryR := strings.Join(parts, " | ")

	// Search in open_positions
	query_open := `
		SELECT 	posts.pid, open_positions.oid, open_positions.title
		FROM posts

		INNER JOIN posts_open_positions ON posts.pid = posts_open_positions.pid
		INNER JOIN open_positions ON posts_open_positions.oid = open_positions.oid

		WHERE
			(to_tsquery($1) @@ to_tsvector(open_positions.title)
			OR SIMILARITY($2, open_positions.title) > 0 )
			AND posts.uid = ANY($3)
		ORDER BY
			NULLIF(ts_rank(to_tsvector(open_positions.title), to_tsquery($1)), 0) DESC,
			SIMILARITY($2, open_positions.title) DESC NULLS LAST
	`
	rows, err := r.db.QueryContext(ctx, query_open, tqueryO, search.SearchOpenPosition, pq.Array(cids))
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	for rows.Next() {
		var pid, oid int64
		var title string
		err = rows.Scan(&pid, &oid, &title)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}

		a := &[]domain.SubSearch{}
		searchResult.OpenPosition[pid] = searchResult.OpenPosition[pid].Append(a, domain.SubSearch{
			Id:    oid,
			Title: title,
		})
	}

	// Search in required_skills
	query_required := `
		SELECT 	posts.pid, required_skills.sid, required_skills.title
		FROM posts

		INNER JOIN posts_required_skills ON posts.pid = posts_required_skills.pid
		INNER JOIN required_skills ON posts_required_skills.sid = required_skills.sid

		WHERE
			(to_tsquery($1) @@ to_tsvector(required_skills.title)
			OR SIMILARITY($2, required_skills.title) > 0 )
			AND posts.uid = ANY($3)
		ORDER BY
			NULLIF(ts_rank(to_tsvector(required_skills.title), to_tsquery($1)), 0) DESC,
			SIMILARITY($2, required_skills.title) DESC NULLS LAST
	`

	rows, err = r.db.QueryContext(ctx, query_required, tqueryR, search.SearchRequiredSkill, pq.Array(cids))
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	for rows.Next() {
		var pid, sid int64
		var title string
		err = rows.Scan(&pid, &sid, &title)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}

		searchResult.RequiredSkill[pid] = domain.SubSearch{
			Id:    sid,
			Title: title,
		}
	}

	// Search in benefits
	query_benefit := `
		SELECT 	posts.pid, benefits.bid, benefits.title
		FROM posts

		INNER JOIN posts_benefits ON posts.pid = posts_benefits.pid
		INNER JOIN benefits ON posts_benefits.bid = benefits.bid

		WHERE
			(to_tsquery($1) @@ to_tsvector(benefits.title)
			OR SIMILARITY($2, benefits.title) > 0 )
			AND posts.uid = ANY($3)
		ORDER BY
			NULLIF(ts_rank(to_tsvector(benefits.title), to_tsquery($1)), 0) DESC,
			SIMILARITY($2, benefits.title) DESC NULLS LAST
	`

	rows, err = r.db.QueryContext(ctx, query_benefit, tqueryB, search.SearchBenefit, pq.Array(cids))
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	for rows.Next() {
		var pid, bid int64
		var title string
		err = rows.Scan(&pid, &bid, &title)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}

		searchResult.Benefits[pid] = domain.SubSearch{
			Id:    bid,
			Title: title,
		}
	}

	// Find posts that apper in all 3 maps
	var postIds []int64
	for pid, _ := range searchResult.OpenPosition {
		if _, ok := searchResult.RequiredSkill[pid]; ok {
			if _, ok := searchResult.Benefits[pid]; ok {
				postIds = append(postIds, pid)
			}
		}
	}

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

func (r *postRepository) UpdatePost(ctx context.Context, postId int64, post *pbv1.Post) error {
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

	// Find and Delete from table `posts_open_positions` and table `open_positions`
	err = r.findAndDeleteOpenPositions(ctx, tx, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Find and Delete from table `posts_required_skills` and table `required_skills`
	err = r.findAndDeleteRequiredSkills(ctx, tx, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Find and Delete from table `posts_benefits` and table `benefits`
	err = r.findAndDeleteBenefits(ctx, tx, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

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

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *postRepository) DeletePost(ctx context.Context, postId int64) error {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Find and Delete from table `posts_open_positions` and table `open_positions`
	err = r.findAndDeleteOpenPositions(ctx, tx, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Find and Delete from table `posts_required_skills` and table `required_skills`
	err = r.findAndDeleteRequiredSkills(ctx, tx, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Find and Delete from table `posts_benefits` and table `benefits`
	err = r.findAndDeleteBenefits(ctx, tx, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from table posts
	query := "DELETE FROM posts WHERE pid = $1"
	_, err = tx.ExecContext(ctx, query, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}
