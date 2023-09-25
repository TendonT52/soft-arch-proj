package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TikhampornSky/go-mail/config"
	"github.com/TikhampornSky/go-mail/service"
	"github.com/memphisdev/memphis.go"
)

const (
	StudentConfirmEmail = "student_confirm_email"
	CompanyApproveEmail = "company_approve_email"
	CompanyRejectEmail  = "company_reject_email"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	conn, err := memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Consumer connection failed: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	consumer, err := conn.CreateConsumer(config.MemphisStationName, config.MemphisConsumer, memphis.PullInterval(1*time.Second))
	if err != nil {
		fmt.Printf("Consumer creation failed: %v", err)
		os.Exit(1)
	}

	fmt.Println("Consumer created")

	snmpService := service.NewSMTPService(&config)
	templateService := service.NewTemplateService(&config, snmpService)

	ch := make(chan int)
	handler := func(msgs []*memphis.Msg, err error, ctx context.Context) {
		if err != nil {
			fmt.Printf("Fetch failed: %v", err)
			return
		}

		for _, msg := range msgs {
			headers := msg.GetHeaders()
			messageType := headers["type"]
			data := msg.Data()
			log.Println("Message type...: ", messageType)
			err := templateService.ProcessEmailRequest(data, messageType)
			if err != nil {
				log.Println("Error while process data: ", err)
			} else {
				msg.Ack()
			}
		}
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "hdr-key", "hdr-value")
	consumer.SetContext(ctx)
	consumer.Consume(handler)
	<-ch
}