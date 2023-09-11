package port

import (
	"bytes"
	pbv1 "github.com/TikhampornSky/go-mail/gen/v1"
)

type TemplateService interface {
	ProcessEmailRequest(rawData []byte, typeMail string) error
}

type SmtpService interface {
	SendEmail(data *pbv1.EmailData, body bytes.Buffer) error
}
