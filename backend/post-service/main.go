package main

import (
	"log"

	"github.com/TikhampornSky/go-post-service/server"
	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/db"
	"github.com/TikhampornSky/go-post-service/repo"
	"github.com/TikhampornSky/go-post-service/service"
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

	postRepo := repo.NewPostRepository(db)
	postService := service.NewPostService(postRepo)

	server.NewServer(config.ServerPort, postService)
}
