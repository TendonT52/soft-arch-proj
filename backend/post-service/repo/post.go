package repo

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
	"github.com/TikhampornSky/go-post-service/utils"
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

func (r *postRepository) CreatePost(ctx context.Context, userId int64, post *pbv1.CreatedPost) (int64, error) {
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
	err = r.insertIntoOpenPositions(ctx, tx, &post.OpenPositions, postId)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into required_skills and posts_required_skills
	err = r.insertIntoRequiredSkills(ctx, tx, &post.RequiredSkills, postId)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into benefits and posts_benefits
	err = r.insertIntoBenefits(ctx, tx, &post.Benefits, postId)
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
	query := "SELECT topic, description, period, how_to, uid, updated_at FROM posts WHERE pid = $1"
	var topic, description, period, howTo string
	var userId, updated_at int64
	err := r.db.QueryRowContext(ctx, query, postId).Scan(&topic, &description, &period, &howTo, &userId, &updated_at)
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
		UpdatedAt: updated_at,
	}

	return post, nil
}

func (r *postRepository) GetPosts(ctx context.Context, search *pbv1.SearchOptions, data *domain.CompanyInfo) ([]*pbv1.Post, error) {
	cids := *data.Ids
	// Create
	parts := strings.Fields(search.SearchBenefit)
	tqueryB := strings.Join(parts, " | ")

	parts = strings.Fields(search.SearchOpenPosition)
	tqueryO := strings.Join(parts, " | ")

	parts = strings.Fields(search.SearchRequiredSkill)
	tqueryR := strings.Join(parts, " | ")

	// // Search in open_positions
	searchOpenPosition, err := r.searchOpenPositions(ctx, tqueryO, search.SearchOpenPosition, &cids)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}

	// Search in required_skills
	searchRequiredSkill, err := r.searchRequiredSkills(ctx, tqueryR, search.SearchRequiredSkill, &cids)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}

	// Search in benefits
	searchBenefit, err := r.searchBenefits(ctx, tqueryB, search.SearchBenefit, &cids)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}

	// Find posts that apper in all 3 maps
	var postIds []int64
	for _, pid := range *searchOpenPosition {
		if utils.IsItemInArray(pid, searchRequiredSkill) && utils.IsItemInArray(pid, searchBenefit) {
			postIds = append(postIds, pid)
		}
	}

	// Find selected posts detail
	var posts []*pbv1.Post
	query := `
			SELECT
			posts.pid,
			uid,
			topic,
			description,
			period,
			how_to,
			posts.updated_at,
			ARRAY_AGG(DISTINCT open_positions.title) AS openPositions,
			ARRAY_AGG(DISTINCT required_skills.title) AS requiredSkills,
			ARRAY_AGG(DISTINCT benefits.title) AS benefits
		FROM
			posts
			INNER JOIN posts_open_positions ON posts.pid = posts_open_positions.pid
			INNER JOIN open_positions ON posts_open_positions.oid = open_positions.oid
			INNER JOIN posts_required_skills ON posts.pid = posts_required_skills.pid
			INNER JOIN required_skills ON posts_required_skills.sid = required_skills.sid
			INNER JOIN posts_benefits ON posts.pid = posts_benefits.pid
			INNER JOIN benefits ON posts_benefits.bid = benefits.bid
		WHERE
			posts.pid = $1
		GROUP BY posts.pid;
	`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt.Close()

	for _, postId := range postIds {
		var pid, ownerId, updated_at int64
		var topic, description, period, howTo string
		var openPositions, requiredSkills, benefits []string
		err = stmt.QueryRowContext(ctx, postId).Scan(&pid, &ownerId, &topic, &description, &period, &howTo, &updated_at, pq.Array(&openPositions), pq.Array(&requiredSkills), pq.Array(&benefits))
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		posts = append(posts, &pbv1.Post{
			PostId:      pid,
			Topic:       topic,
			Description: description,
			Period:      period,
			HowTo:       howTo,
			Owner: &pbv1.PostOwner{
				Id:   ownerId,
				Name: data.Profiles[ownerId].Name,
			},
			OpenPositions:  openPositions,
			RequiredSkills: requiredSkills,
			Benefits:       benefits,
			UpdatedAt:      updated_at,
		})
	}

	return posts, nil
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

func (r *postRepository) UpdatePost(ctx context.Context, postId int64, post *pbv1.UpdatedPost) error {
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

	err = r.updateOpenPosition(ctx, tx, post.OpenPositions, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	err = r.updateRequiredSkill(ctx, tx, post.RequiredSkills, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	err = r.updateBenefit(ctx, tx, post.Benefits, postId)
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

	// Select open_positions, required_skills, benefits
	query := "SELECT title FROM open_positions WHERE oid IN (SELECT oid FROM posts_open_positions WHERE pid = $1)"
	rows, err := tx.QueryContext(ctx, query, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	var openPositions []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		openPositions = append(openPositions, title)
	}

	query = "SELECT title FROM required_skills WHERE sid IN (SELECT sid FROM posts_required_skills WHERE pid = $1)"
	rows, err = tx.QueryContext(ctx, query, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()
	var requiredSkills []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		requiredSkills = append(requiredSkills, title)
	}

	query = "SELECT title FROM benefits WHERE bid IN (SELECT bid FROM posts_benefits WHERE pid = $1)"
	rows, err = tx.QueryContext(ctx, query, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	var benefits []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			tx.Rollback()
			return domain.ErrInternal.From(err.Error(), err)
		}
		benefits = append(benefits, title)
	}

	// Find and Delete from table `posts_open_positions` and table `open_positions`
	err = r.deleteOpenPositions(ctx, tx, &openPositions, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Find and Delete from table `posts_required_skills` and table `required_skills`
	err = r.deleteRequiredSkills(ctx, tx, &requiredSkills, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Find and Delete from table `posts_benefits` and table `benefits`
	err = r.deleteBenefits(ctx, tx, &benefits, postId)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from table posts
	query = "DELETE FROM posts WHERE pid = $1"
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

func (r *postRepository) GetOpenPositions(ctx context.Context, search string) ([]string, error) {
	parts := strings.Fields(search)
	tquery := strings.Join(parts, " | ")

	var rows *sql.Rows
	var err error
	if search == "" {
		query := "SELECT title FROM open_positions ORDER BY title ASC"
		rows, err = r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	} else {
		query :=
			`
		SELECT
			title
		FROM
			open_positions,
			to_tsvector(title) document,
			to_tsquery($1) query,
			NULLIF(ts_rank(to_tsvector(title), query), 0) rank,
			SIMILARITY ($2, title) similarity
		WHERE
			query @@ document
			OR similarity > 0
		ORDER BY
			rank DESC NULLS LAST,
			similarity DESC NULLS LAST
		`
		rows, err = r.db.QueryContext(ctx, query, tquery, search)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
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
	return openPositions, nil
}

func (r *postRepository) GetRequiredSkills(ctx context.Context, search string) ([]string, error) {
	parts := strings.Fields(search)
	tquery := strings.Join(parts, " | ")

	var rows *sql.Rows
	var err error
	if search == "" {
		query := "SELECT title FROM required_skills ORDER BY title ASC"
		rows, err = r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	} else {
		query :=
			`
		SELECT
			title
		FROM
			required_skills,
			to_tsvector(title) document,
			to_tsquery($1) query,
			NULLIF(ts_rank(to_tsvector(title), query), 0) rank,
			SIMILARITY ($2, title) similarity
		WHERE
			query @@ document
			OR similarity > 0
		ORDER BY
			rank DESC NULLS LAST,
			similarity DESC NULLS LAST
		`
		rows, err = r.db.QueryContext(ctx, query, tquery, search)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
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
	return requiredSkills, nil
}

func (r *postRepository) GetBenefits(ctx context.Context, search string) ([]string, error) {
	parts := strings.Fields(search)
	tquery := strings.Join(parts, " | ")

	var rows *sql.Rows
	var err error
	if search == "" {
		query := "SELECT title FROM benefits ORDER BY title ASC"
		rows, err = r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	} else {
		query :=
			`
		SELECT
			title
		FROM
			benefits,
			to_tsvector(title) document,
			to_tsquery($1) query,
			NULLIF(ts_rank(to_tsvector(title), query), 0) rank,
			SIMILARITY ($2, title) similarity
		WHERE
			query @@ document
			OR similarity > 0
		ORDER BY
			rank DESC NULLS LAST,
			similarity DESC NULLS LAST
		`
		rows, err = r.db.QueryContext(ctx, query, tquery, search)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
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
	return benefits, nil
}

func (r *postRepository) GetMyPosts(ctx context.Context, userId int64) ([]*pbv1.Post, error) {
	query := "SELECT pid, topic, description, period, how_to, updated_at FROM posts WHERE uid = $1 ORDER BY topic ASC"
	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	var posts []*pbv1.Post
	for rows.Next() {
		var pid, updated_at int64
		var topic, description, period, howTo string
		err = rows.Scan(&pid, &topic, &description, &period, &howTo, &updated_at)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		posts = append(posts, &pbv1.Post{
			PostId:      pid,
			Topic:       topic,
			Description: description,
			Period:      period,
			HowTo:       howTo,
			UpdatedAt:   updated_at,
		})
	}

	// Get open_positions
	prepare0 := "SELECT title FROM open_positions WHERE oid IN (SELECT oid FROM posts_open_positions WHERE pid = $1)"
	stmt0, err := r.db.PrepareContext(ctx, prepare0)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt0.Close()

	for _, post := range posts {
		rows, err = stmt0.QueryContext(ctx, post.PostId)
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
		post.OpenPositions = openPositions
	}

	// Get required_skills
	prepare1 := "SELECT title FROM required_skills WHERE sid IN (SELECT sid FROM posts_required_skills WHERE pid = $1)"
	stmt1, err := r.db.PrepareContext(ctx, prepare1)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt1.Close()

	for _, post := range posts {
		rows, err = stmt1.QueryContext(ctx, post.PostId)
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
		post.RequiredSkills = requiredSkills
	}

	// Get benefits
	prepare2 := "SELECT title FROM benefits WHERE bid IN (SELECT bid FROM posts_benefits WHERE pid = $1)"
	stmt2, err := r.db.PrepareContext(ctx, prepare2)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt2.Close()

	for _, post := range posts {
		rows, err = stmt2.QueryContext(ctx, post.PostId)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}

		var benefits []string
		for rows.Next() {
			var title string
			err = rows.Scan(&title)
			if err != nil {
				return nil, domain.ErrInternal.From(err.Error(), err)
			}
			benefits = append(benefits, title)
		}

		post.Benefits = benefits
	}

	return posts, nil
}
