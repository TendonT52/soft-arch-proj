package e2e

import (
	"fmt"
	"log"
	"testing"

	"github.com/TikhampornSky/go-mail/config"
	"github.com/TikhampornSky/go-mail/domain"
	"github.com/TikhampornSky/go-mail/service"
	"github.com/memphisdev/memphis.go"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	pbv1 "github.com/TikhampornSky/go-mail/gen/v1"
)

func TestStudentVerification(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	require.NoError(t, err)

	conn, err := memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Consumer connection failed: %v", err)
	}
	require.NoError(t, err)
	defer conn.Close()

	emailData := pbv1.EmailData{
		URL:     "http://localhost:3000/verify/123456",
		Subject: "Verify your email",
		Name:    "Tikhamporn",
		Email:   "mock_student@chula.ac.th",
	}

	data, err := proto.Marshal(&emailData)
	require.NoError(t, err)

	createMockProducer(domain.StudentConfirmEmail, data)
	c := createConsumer(service.NewTemplateService(&config, service.NewMockSMTPService(&config)))
	require.Nil(t, c)
}

func TestCompanyApprove(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	require.NoError(t, err)

	conn, err := memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Consumer connection failed: %v", err)
	}
	require.NoError(t, err)
	defer conn.Close()

	emailData := pbv1.EmailData{
		Subject: "Approve Company",
		Name:    "Tikhamporn",
		Email:   "company1@gmail.comm",
	}

	data, err := proto.Marshal(&emailData)
	require.NoError(t, err)

	createMockProducer(domain.CompanyApproveEmail, data)
	c := createConsumer(service.NewTemplateService(&config, service.NewMockSMTPService(&config)))
	require.Nil(t, c)
}

func TestCompanyReject(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	require.NoError(t, err)

	conn, err := memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Consumer connection failed: %v", err)
	}
	require.NoError(t, err)
	defer conn.Close()

	emailData := pbv1.EmailData{
		Subject: "Reject Company",
		Name:    "Tepsut",
		Email:   "company2@gmail.comm",
	}

	data, err := proto.Marshal(&emailData)
	require.NoError(t, err)

	createMockProducer(domain.CompanyRejectEmail, data)
	c := createConsumer(service.NewTemplateService(&config, service.NewMockSMTPService(&config)))
	require.Nil(t, c)
}

func TestUnknowEmailType(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	require.NoError(t, err)

	conn, err := memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Consumer connection failed: %v", err)
	}
	require.NoError(t, err)
	defer conn.Close()

	emailData := pbv1.EmailData{
		Subject: "Approve Company",
		Name:    "Tikhamporn",
		Email:   "company3@gmail.comm",
	}

	data, err := proto.Marshal(&emailData)
	require.NoError(t, err)

	createMockProducer("UnknownType", data)
	c := createConsumer(service.NewTemplateService(&config, service.NewMockSMTPService(&config)))
	require.ErrorIs(t, c, domain.ErrKindUnknown)
}
