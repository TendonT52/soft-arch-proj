package e2e

import (
	"log"
	"os"
	"testing"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/db"
	"github.com/TikhampornSky/go-post-service/repo"
	"github.com/TikhampornSky/go-post-service/server"
	"github.com/TikhampornSky/go-post-service/service"
)

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("..")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db, err := db.NewDatabase(config)
	if err != nil {
		log.Fatalf("Something went wrong. Could not connect to the database. %s", err)
	}

	postRepo := repo.NewPostRepository(db.GetPostgresqlDB())
	tokenService := service.NewTokenTestService()
	postService := service.NewPostService(postRepo, tokenService)

	// gRPC Zone
	go server.NewServer(config.ServerPort, postService)

	code := m.Run()
	os.Exit(code)
}
