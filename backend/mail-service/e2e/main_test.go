package e2e

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/TikhampornSky/go-mail/config"
	"github.com/TikhampornSky/go-mail/port"
	"github.com/memphisdev/memphis.go"
)

var (
	MockProducer *memphis.Producer
	MockConsumer *memphis.Consumer
	Conn         *memphis.Conn
)

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	Conn, err = memphis.Connect(config.MemphisHostName, config.MemphisApplicationUser, memphis.Password(config.MemphisPassword), memphis.AccountId(config.MemphisAccountID))
	if err != nil {
		fmt.Printf("Consumer connection failed: %v", err)
		os.Exit(1)
	}
	defer Conn.Close()

	err = clearUnprocessedMessage()
	if err != nil {
		fmt.Printf("Clear unprocessed message failed: %v", err)
		os.Exit(1)
	}

	MockProducer, err = Conn.CreateProducer(config.MemphisStationNameTest, config.MemphisProducerTest)
	if err != nil {
		fmt.Printf("Producer failed: %v", err)
		os.Exit(1)
	}

	MockConsumer, err = Conn.CreateConsumer(config.MemphisStationNameTest, config.MemphisConsumer, memphis.PullInterval(1*time.Second))
	if err != nil {
		fmt.Printf("Consumer creation failed: %v", err)
		os.Exit(1)
	}

	os.Chdir("../")
	code := m.Run()
	os.Exit(code)
}

func createConsumer(templateService port.TemplateService) error {
	ch := make(chan int)
	var errc []error
	handler := func(msgs []*memphis.Msg, err error, ctx context.Context) {
		if err != nil {
			fmt.Printf("Fetch failed: %v", err)
		}

		for _, msg := range msgs {
			headers := msg.GetHeaders()
			messageType := headers["type"]
			data := msg.Data()
			fmt.Println("Message type...: ", messageType)
			err = templateService.ProcessEmailRequest(data, messageType)
			if err != nil {
				errc = append(errc, err)
			} else {
				msg.Ack()
			}
		}

		ch <- 1
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "hdr-key", "hdr-value")
	MockConsumer.SetContext(ctx)
	MockConsumer.Consume(handler)
	<-ch

	if len(errc) != 0 {
		return errc[0]
	}

	return nil
}

func createMockProducer(typeMail string, jsonData string) {
	hdrs := memphis.Headers{}
	hdrs.New()
	err := hdrs.Add("type", typeMail)
	if err != nil {
		fmt.Printf("Header failed: %v", err)
		os.Exit(1)
	}

	err = MockProducer.Produce([]byte(jsonData), memphis.MsgHeaders(hdrs))
	if err != nil {
		fmt.Printf("Produce failed: %v", err)
		os.Exit(1)
	}
}

func clearUnprocessedMessage() error {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	consumer, err := Conn.CreateConsumer(config.MemphisStationNameTest, config.MemphisConsumer, memphis.PullInterval(1*time.Second))
	if err != nil {
		fmt.Printf("Consumer creation failed: %v", err)
		return err
	}

	fmt.Println("CLEAR Consumer created")
	ok := true
	for {
		msgs, err := consumer.Fetch(100, true)
		if err != nil { // no more messages
			ok = false
		}

		if !ok {
			break
		}

		for _, msg := range msgs {
			msg.Ack()
		}
	}

	return nil
}
