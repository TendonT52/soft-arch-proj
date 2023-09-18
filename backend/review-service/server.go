package main

import (
	"JinnnDamanee/review-service/config"
	"JinnnDamanee/review-service/db"
	"JinnnDamanee/review-service/httpServer"
	"JinnnDamanee/review-service/internal/handler"
	"JinnnDamanee/review-service/internal/repo"
	"JinnnDamanee/review-service/internal/service"

	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}

type Review struct {
	gorm.Model
	text string
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config: ", err)
	}

	db, err := db.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	reviewRepo := repo.NewReviewRepository(db.Gorm)
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewService)
	s := httpServer.NewHTTPServer(reviewService, reviewHandler)
	s.InitRouter()

	if err != nil {
		log.Fatal(err)
	}

	s.SetUpShutdown()
}
