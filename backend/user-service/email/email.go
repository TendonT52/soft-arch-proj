package email

import (
	"fmt"

	"github.com/TikhampornSky/go-auth-verifiedMail/initializers"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/memphisdev/memphis.go"
)

func SendEmail(conn port.MemphisPort, typeMail string, jsonData []byte) error {
	config, err := initializers.LoadConfig("../")
	if err != nil {
		fmt.Printf("Config failed: %v", err)
		return err
	}

	p, err := conn.CreateProducer(conn.GetStationName(), config.MemphisProducer)
	if err != nil {
		fmt.Printf("Producer failed: %v", err)
		return err
	}

	hdrs := memphis.Headers{}
	hdrs.New()
	err = hdrs.Add("type", typeMail)
	if err != nil {
		fmt.Printf("Header failed: %v", err)
		return err
	}

	err = p.Produce([]byte(jsonData), memphis.MsgHeaders(hdrs))
	if err != nil {
		fmt.Printf("Produce failed: %v", err)
		return err
	}
	return nil
}

func IsChulaStudentEmail(email string) bool {
	return email[len(email)-20:] == "@student.chula.ac.th"
}

func GetStudentIDFromEmail(email string) string {
	return email[:10]
}
