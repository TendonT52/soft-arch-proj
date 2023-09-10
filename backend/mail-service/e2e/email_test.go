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

	err = clearUnprocessedMessage()
	require.NoError(t, err)

	createMockProducer(domain.StudentConfirmEmail, `{"url": "http://localhost:3000/verify/123456", "subject": "Verify your email", "name": "Tikhamporn", "email": "mock_student@chula.ac.th"}`)
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

	err = clearUnprocessedMessage()
	require.NoError(t, err)

	createMockProducer(domain.CompanyApproveEmail, `{"url": "http://localhost:3000/verify/123456", "subject": "Approve Company", "name": "Mock Company", "email": "company@gmail.com"}`)
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

	err = clearUnprocessedMessage()
	require.NoError(t, err)

	createMockProducer(domain.CompanyRejectEmail, `{"url": "http://localhost:3000/verify/123456", "subject": "Reject Company", "name": "Mock Company2", "email": "company2@gmail.com"}`)
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

	err = clearUnprocessedMessage()
	require.NoError(t, err)

	createMockProducer("UnknownType", `{"url": "http://localhost:3000/verify/123456", "subject": "Reject Company", "name": "Mock Company2", "email": "company2@gmail.com"}`)
	c := createConsumer(service.NewTemplateService(&config, service.NewMockSMTPService(&config)))
	require.ErrorIs(t, c, domain.ErrKindUnknown)
}
