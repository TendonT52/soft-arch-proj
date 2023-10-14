package tools

import (
	"database/sql"
	"fmt"

	"github.com/JinnnDamanee/review-service/config"
	_ "github.com/lib/pq"
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

func DeleteAllRecords() error {
	db, err := createDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM reviews")
	if err != nil {
		return err
	}

	return nil
}
