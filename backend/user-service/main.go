package main

import (
	"log"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/db"
	"github.com/TikhampornSky/go-auth-verifiedMail/email"
	"github.com/TikhampornSky/go-auth-verifiedMail/repo"
	"github.com/TikhampornSky/go-auth-verifiedMail/server"
	"github.com/TikhampornSky/go-auth-verifiedMail/service"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db, err := db.NewDatabase(&config)
	if err != nil {
		log.Fatalf("Something went wrong. Could not connect to the database. %s", err)
	}

	memphisConn := email.InitMemphisConnection()
	defer memphisConn.Close()

	memphis := email.NewMemphis(memphisConn, config.MemphisStationName)

	userRepo := repo.NewUserRepository(db.GetPostgresqlDB(), db.GetRedisDB())
	timeService := service.NewRealTimeProvider()
	authService := service.NewAuthService(userRepo, memphis, timeService)
	userService := service.NewUserService(userRepo, memphis)

	// gRPC Zone
	server.NewServer(config.ServerPort, authService, userService)
}
