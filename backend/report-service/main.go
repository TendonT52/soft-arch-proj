package main

import (
	"log"

	"github.com/TikhampornSky/report-service/repo"
	"github.com/TikhampornSky/report-service/service"
	"github.com/TikhampornSky/report-service/config"
	"github.com/TikhampornSky/report-service/db"
	"github.com/TikhampornSky/report-service/server"
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
