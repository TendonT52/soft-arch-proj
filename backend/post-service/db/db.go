package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/TikhampornSky/go-post-service/config"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sql.DB
}

// BeginTx implements repo.DBTX.
func (*Database) BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error) {
	panic("unimplemented")
}

// ExecContext implements repo.DBTX.
func (*Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("unimplemented")
}

// PrepareContext implements repo.DBTX.
func (*Database) PrepareContext(context.Context, string) (*sql.Stmt, error) {
	panic("unimplemented")
}

// QueryContext implements repo.DBTX.
func (*Database) QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error) {
	panic("unimplemented")
}

// QueryRowContext implements repo.DBTX.
func (*Database) QueryRowContext(context.Context, string, ...interface{}) *sql.Row {
	panic("unimplemented")
}

func NewDatabase(config *config.Config) (*Database, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName, "disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to the postgresql database")

	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	return nil, err
	// }
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://"+config.MigrationPath,
	// 	"postgres", driver)

	// if m == nil {
	// 	return nil, err
	// }
	// err = m.Up()
	// if err != nil && err != migrate.ErrNoChange {
	// 	return nil, err
	// }
	// log.Println("Successfully applied migrations")

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetPostgresqlDB() *sql.DB {
	return d.db
}
