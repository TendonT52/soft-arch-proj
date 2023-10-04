package email

import (
	"fmt"
	"regexp"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/asaskevich/govalidator"
	"github.com/memphisdev/memphis.go"
	"google.golang.org/protobuf/proto"
)

func SendEmail(conn port.MemphisPort, typeMail, url, subject, name, email string) error {
	config, err := config.LoadConfig("../")
	if err != nil {
		fmt.Printf("Config failed: %v", err)
		return err
	}

	// Send Email
	emailData := pbv1.EmailData{
		URL:     url,
		Subject: subject,
		Name:    name,
		Email:   email,
	}

	data, err := proto.Marshal(&emailData)
	if err != nil {
		fmt.Println("Error:", err)
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

	err = p.Produce([]byte(data), memphis.MsgHeaders(hdrs))
	if err != nil {
		fmt.Printf("Produce failed: %v", err)
		return err
	}
	return nil
}

func IsChulaStudentEmail(email string) bool {
	customEmailPattern := `^\d{10}@student\.chula\.ac\.th$`
	regex := regexp.MustCompile(customEmailPattern)
	return regex.MatchString(email)
}

func GetStudentIDFromEmail(email string) string {
	return email[:10]
}

func IsCorrectEmailFormat(email string) bool {
	return govalidator.IsEmail(email)
}
