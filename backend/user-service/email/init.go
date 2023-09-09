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
	stationName string
}

func NewMemphis(conn *memphis.Conn, stationName string) *Memphis {
	return &Memphis{
		conn: conn,
		stationName: stationName,
	}
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

func (m *Memphis) CreateProducer(stationName string, name string, opts ...memphis.ProducerOpt) (*memphis.Producer, error) {
	return m.conn.CreateProducer(m.stationName, name, opts...)
}

func (m *Memphis) GetStationName() string {
	return m.stationName
}
