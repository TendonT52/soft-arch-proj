package email

import (
	"fmt"
	"log"
	"os"

	"github.com/TikhampornSky/go-auth-verifiedMail/initializers"
	"github.com/memphisdev/memphis.go"
)

type Memphis struct {
	conn *memphis.Conn
}

func InitMemphisConnection() *memphis.Conn {
	config, err := initializers.LoadConfig("../")
	if err != nil {
		fmt.Printf("Config failed: %v", err)
		os.Exit(1)
	}

	conn, err := memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Connection failed: %v", err)
		os.Exit(1)
	}
	log.Println("Memphis connection established")

	return conn
}
