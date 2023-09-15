package main

import (
	"JinnnDamanee/review-service/db"
	"JinnnDamanee/review-service/httpServer"
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
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	reviewRepo := repo.NewReviewRepository(db.Gorm)
	reviewService := service.NewReviewService(reviewRepo)
	s := httpServer.NewHTTPServer(reviewService)
	s.InitRouter()

	if err != nil {
		log.Fatal(err)
	}

	s.SetUpShutdown()
}
