package service

import (
	"bytes"
	"crypto/tls"
	"log"

	"github.com/TikhampornSky/go-mail/config"
	"github.com/TikhampornSky/go-mail/domain"
	"github.com/TikhampornSky/go-mail/port"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type smtpService struct {
	Config *config.Config
}

func NewSMTPService(config *config.Config) port.SmtpService {
	return &smtpService{Config: config}
}

func (s *smtpService) SendEmail(data *domain.EmailData, body bytes.Buffer) error {
	config, err := config.LoadConfig("../")

	if err != nil {
		log.Fatal("could not load config", err)
		return err
	}

	from := config.EmailFrom
	smtpPass := config.SMTPPass
	to := data.Email
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort
	smtpFrom := config.EmailFrom

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpFrom, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
		return err
	}

	return nil
}
