package port

import (
	"bytes"

	"github.com/TikhampornSky/go-mail/domain"
)

type TemplateService interface {
	ProcessEmailRequest(rawData []byte, typeMail string) error
}

type SmtpService interface {
	SendEmail(data *domain.EmailData, body bytes.Buffer) error
}
