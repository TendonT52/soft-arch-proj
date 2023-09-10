package email

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/TikhampornSky/go-mail/config"
	"github.com/TikhampornSky/go-mail/domain"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

func ParseTemplateDir(dir string) (*template.Template, error) {
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

func SendStudentEmail(data *domain.EmailData, typeMail string) error {
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

	var body bytes.Buffer

	if typeMail == domain.StudentConfirmEmail {
		log.Println("Send email to student")
		template, err := ParseTemplateDir("templates")
		if err != nil {
			log.Fatal("Could not parse template", err)
			return err
		}
		template.ExecuteTemplate(&body, "verificationCode.html", &data)
	} else if typeMail == domain.CompanyApproveEmail {
		log.Println("Send email to company (Approve)")
		template, err := ParseTemplateDir("templates-company-approve")
		if err != nil {
			log.Fatal("Could not parse template", err)
			return err
		}
		template.ExecuteTemplate(&body, "verity-approve.html", &data)
	} else if typeMail == domain.CompanyRejectEmail {
		log.Println("Send email to company (Reject)")
		template, err := ParseTemplateDir("templates-company-reject")
		if err != nil {
			log.Fatal("Could not parse template", err)
			return err
		}
		template.ExecuteTemplate(&body, "verify-reject.html", &data)
	} else {
		log.Fatal("Unknown message type")
		return err
	}

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
