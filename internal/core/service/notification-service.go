package service

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"notification-api/internal/core/domain"
	"notification-api/internal/ports"
	"os"
	"os/signal"
	"syscall"
)

type NotificationService struct {
	consumer   ports.Consumer
	httpClient ports.HttpClient
}

func NewNotificationService(consumer ports.Consumer, httpClient ports.HttpClient) *NotificationService {
	return &NotificationService{
		consumer:   consumer,
		httpClient: httpClient,
	}
}

func (ns *NotificationService) ProcessNotifications() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Println("Warning: Error loading .env file. Using environment variables.")
	}

	var host = os.Getenv("BASE_URI_NOTIFICATION")
	var path = os.Getenv("PATH_MAIL_NOTIFICATION")
	var url = buildUrl(host, path)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		log.Println("Listening to topic...")
		select {
		case <-sigChan:
			log.Println("Received shutdown signal, exiting...")
			return
		default:
			message := ns.consumer.Consume()
			var notification domain.Notification
			err := json.Unmarshal([]byte(message), &notification)
			if err != nil {
				log.Printf("error to deserialize message: %v", err)
				continue
			}

			_, err = ns.httpClient.POST(url, message, nil)
			if err != nil {
				log.Printf("error: %v", err)
				// send bucket or retry process
			}

			// send message to topic FinishRetryProcessTopic, to finish retry process.
		}
	}
}

func buildUrl(host string, path string) string {
	return fmt.Sprintf("%s%s", host, path)
}
