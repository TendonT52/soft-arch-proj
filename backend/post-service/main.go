package main

import (
	"log"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/db"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db, err := db.NewDatabase(config)
	if err != nil {
		log.Fatalf("Something went wrong. Could not connect to the database. %s", err)
	}
	defer db.Close()
}
