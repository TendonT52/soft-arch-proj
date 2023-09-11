package service

import (
	"bytes"

	"github.com/TikhampornSky/go-mail/config"
	pbv1 "github.com/TikhampornSky/go-mail/gen/v1"
	"github.com/TikhampornSky/go-mail/port"
)

type mockSmtpService struct {
	Config *config.Config
}

func NewMockSMTPService(config *config.Config) port.SmtpService {
	return &mockSmtpService{Config: config}
}

func (s *mockSmtpService) SendEmail(data *pbv1.EmailData, body bytes.Buffer) error {
	return nil
}
