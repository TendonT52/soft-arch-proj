package tools

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/domain"
)

func createDBConnection() (*sql.DB, error) {
	config, err := config.LoadConfig("..")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName, "disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DeleteAllPosts() error {
	db, err := createDBConnection()
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
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
