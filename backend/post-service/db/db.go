package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/TikhampornSky/go-post-service/config"
)

type Database struct {
	db *sql.DB
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
