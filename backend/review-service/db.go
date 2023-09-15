package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

func NewDatabase() (*Database, error) {

	connFormat := "user=%s password=%s dbname=%s host=%s port=%s sslmode=disable"
	conn := fmt.Sprintf(connFormat, "root", "password", "review", "localhost", "5432")

	// sqlDB, err := sql.Open("pgx", conn)
	// if err != nil {
	// 	return nil, err
	// }
	gormDB, err := gorm.Open(
		postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Print("Connected to database")

	return &Database{Gorm: gormDB}, nil
}

func (d *Database) Close() {
	sqlDB, err := d.Gorm.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

func (d *Database) GetDB() *sql.DB {
	sqlDB, _ := d.Gorm.DB()
	return sqlDB
}
