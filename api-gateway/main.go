package main

import (
	"log"

	"github.com/tendont52/api-gateway/config"
	"github.com/tendont52/api-gateway/gateway"
)

func main() {
	log.Println("Loading config...")
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	log.Println("Config loaded")
	err = gateway.Serve(conf)
	if err != nil {
		log.Fatalf("cannot start gateway: %v", err)
	}
}
