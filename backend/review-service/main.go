package main

import (
	"log"

	"github.com/JinnnDamanee/review-service/config"
	"github.com/JinnnDamanee/review-service/db"
	"github.com/JinnnDamanee/review-service/repo"
	"github.com/JinnnDamanee/review-service/server"
	"github.com/JinnnDamanee/review-service/service"
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

	reportRepo := repo.NewReportRepository(db.GetPostgresqlDB())
	reportService := service.NewReportService(reportRepo)

	server.NewServer(config.ServerPort, reportService)
}
