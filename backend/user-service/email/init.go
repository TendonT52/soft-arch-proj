package email

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/memphisdev/memphis.go"
)

type Memphis struct {
	conn        *memphis.Conn
	stationName string
}

func NewMemphis(conn *memphis.Conn, stationName string) *Memphis {
	return &Memphis{
		conn:        conn,
		stationName: stationName,
	}
}

func InitMemphisConnection() *memphis.Conn {
	config, err := config.LoadConfig("../")
	if err != nil {
		fmt.Printf("Config failed: %v", err)
		os.Exit(1)
	}

	conn, err := memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Connection failed: %v", err)
		os.Exit(1)
	}
	log.Println("Successfully connected to the Memphis server")

	_, err = conn.CreateStation(config.MemphisStationName, 
		memphis.RetentionVal(config.MemphisRetentionVal), 
		memphis.Replicas(config.MemphisReplicas), 
		memphis.IdempotencyWindow(time.Duration(config.MemphisIdempotency) * time.Minute), 
		memphis.PartitionsNumber(config.MemphisPartitions),
	)
	if err != nil {
		fmt.Printf("Create Station failed: %v", err)
		os.Exit(1)
	}

	_, err = conn.CreateStation(config.MemphisStationNameTest, 
		memphis.RetentionVal(config.MemphisRetentionValTest), 
		memphis.Replicas(config.MemphisReplicas), 
		memphis.IdempotencyWindow(time.Duration(config.MemphisIdempotency) * time.Minute), 
		memphis.PartitionsNumber(config.MemphisPartitions),
	)
	if err != nil {
		fmt.Printf("Create Station for test failed: %v", err)
		os.Exit(1)
	}
	log.Println("Successfully created the Memphis station")

	return conn
}

func (m *Memphis) CreateProducer(stationName string, name string, opts ...memphis.ProducerOpt) (*memphis.Producer, error) {
	return m.conn.CreateProducer(m.stationName, name, opts...)
}

func (m *Memphis) GetStationName() string {
	return m.stationName
}
