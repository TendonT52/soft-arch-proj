package repo

import (
	"context"
	"database/sql"
	"strings"
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

func (r *postRepository) GetPosts(ctx context.Context, search *pbv1.SearchOptions, data *domain.CompanyInfo) ([]*pbv1.Post, error) {
	cids := *data.Ids
	var searchResult *domain.IndividualSearchResult = &domain.IndividualSearchResult{
		OpenPositions:  make(map[int64](*[]string)),
		RequiredSkills: make(map[int64](*[]string)),
		Benefits:       make(map[int64](*[]string)),
	}

	// Create
	parts := strings.Fields(search.SearchBenefit)
	tqueryB := strings.Join(parts, " | ")

	parts = strings.Fields(search.SearchOpenPosition)
	tqueryO := strings.Join(parts, " | ")

	parts = strings.Fields(search.SearchRequiredSkill)
	tqueryR := strings.Join(parts, " | ")

	// // Search in open_positions
	searchOpenPosition, orders, err := r.searchOpenPositions(ctx, tqueryO, search.SearchOpenPosition, &cids)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	searchResult.OpenPositions = searchOpenPosition

	// Search in required_skills
	searchRequiredSkill, err := r.searchRequiredSkills(ctx, tqueryR, search.SearchRequiredSkill, &cids)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	searchResult.RequiredSkills = searchRequiredSkill

	// Search in benefits
	searchBenefit, err := r.searchBenefits(ctx, tqueryB, search.SearchBenefit, &cids)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	searchResult.Benefits = searchBenefit

	// Find posts that apper in all 3 maps
	var summary []domain.SummarySearchResult
	for _, pid := range *orders {
		if _, ok := searchResult.OpenPositions[pid]; ok {
			if _, ok := searchResult.RequiredSkills[pid]; ok {
				if _, ok := searchResult.Benefits[pid]; ok {
					summary = append(summary, domain.SummarySearchResult{
						Pid:           pid,
						OpenPosition:  searchResult.OpenPositions[pid],
						RequiredSkill: searchResult.RequiredSkills[pid],
						Benefits:      searchResult.Benefits[pid],
					})
				}
			}
		}
	}

	// Find selected posts detail
	var posts []*pbv1.Post
	query := "SELECT pid, uid, topic, description, period, how_to, updated_at FROM posts WHERE pid = $1"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer stmt.Close()

	for _, s := range summary {
		var pid, ownerId, updated_at int64
		var topic, description, period, howTo string
		err = stmt.QueryRowContext(ctx, s.Pid).Scan(&pid, &ownerId, &topic, &description, &period, &howTo, &updated_at)
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
			OpenPositions:  *s.OpenPosition,
			RequiredSkills: *s.RequiredSkill,
			Benefits:       *s.Benefits,
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

func (r *postRepository) DeleteAllPosts(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	query := "DELETE FROM posts_open_positions"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	query = "DELETE FROM posts_required_skills"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	query = "DELETE FROM posts_benefits"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	query = "DELETE FROM open_positions"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	query = "DELETE FROM required_skills"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}
	query = "DELETE FROM benefits"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	query = "DELETE FROM posts"
	_, err = tx.ExecContext(ctx, query)
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
