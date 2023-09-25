package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/TikhampornSky/go-mail/config"
	"github.com/TikhampornSky/go-mail/domain"
	"github.com/TikhampornSky/go-mail/port"
	pbv1 "github.com/TikhampornSky/go-mail/gen/v1"
	"google.golang.org/protobuf/proto"
)

type templateService struct {
	Config      *config.Config
	SnmpService port.SmtpService
}

func NewTemplateService(config *config.Config, snmp port.SmtpService) port.TemplateService {
	return &templateService{
		Config:      config,
		SnmpService: snmp,
	}
}

func parseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func (e *templateService) ProcessEmailRequest(rawData []byte, typeMail string) error {
	var data pbv1.EmailData
	err := proto.Unmarshal(rawData, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	var body bytes.Buffer

	if typeMail == domain.StudentConfirmEmail {
		log.Println("Send email to student")
		template, err := parseTemplateDir("templates")
		if err != nil {
			log.Fatal("Could not parse template ", err)
			return err
		}
		template.ExecuteTemplate(&body, "verificationCode.html", &data)
	} else if typeMail == domain.CompanyApproveEmail {
		log.Println("Send email to company (Approve)")
		template, err := parseTemplateDir("templates-company-approve")
		if err != nil {
			log.Fatal("Could not parse template ", err)
			return err
		}
		template.ExecuteTemplate(&body, "verity-approve.html", &data)
	} else if typeMail == domain.CompanyRejectEmail {
		log.Println("Send email to company (Reject)")
		template, err := parseTemplateDir("templates-company-reject")
		if err != nil {
			log.Fatal("Could not parse template ", err)
			return err
		}
		template.ExecuteTemplate(&body, "verify-reject.html", &data)
	} else {
		fmt.Println("Unknown message type")
		return domain.ErrKindUnknown
	}

	err = e.SnmpService.SendEmail(&data, body)
	if err != nil {
		fmt.Println("Could not send email: ", err)
		return err
	}

	return nil
}
