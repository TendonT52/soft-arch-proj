package service

import (
	"bytes"

	"github.com/TikhampornSky/go-mail/config"
	"github.com/TikhampornSky/go-mail/domain"
	"github.com/TikhampornSky/go-mail/port"
)

type mockSmtpService struct {
	Config *config.Config
}

func NewMockSMTPService(config *config.Config) port.SmtpService {
	return &mockSmtpService{Config: config}
}

func (s *mockSmtpService) SendEmail(data *domain.EmailData, body bytes.Buffer) error {
	return nil
}
