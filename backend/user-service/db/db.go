package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	db    *sql.DB
	redis *redis.Client
}

func NewDatabase(config *config.Config) (*Database, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName, "disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to the postgresql database")

	redis := redis.NewClient(&redis.Options{
		Addr:     config.REDISHost + ":" + config.REDISPort,
		Password: config.REDISPassword,
		DB:       config.REDISDB,
	})
	log.Println("Successfully connected to the redis database")

	return &Database{db: db, redis: redis}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetPostgresqlDB() *sql.DB {
	return d.db
}

func (d *Database) GetRedisDB() *redis.Client {
	return d.redis
}
