package service

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"notification-api/internal/core/domain"
	"notification-api/internal/ports"
	"os"
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
		log.Fatalf("Error loading .env file")
	}

	var host = os.Getenv("BASE_URI_NOTIFICATION")
	var path = os.Getenv("PATH_MAIL_NOTIFICATION")
	var url = buildUrl(host, path)
	var notification domain.Notification

	message := ns.consumer.Consume()
	err := json.Unmarshal([]byte(message), &notification)
	if err != nil {
		log.Fatalf("error to deserialize message: %v", err)
	}

	_, err = ns.httpClient.POST(url, message, nil)
	if err != nil {
		log.Fatalf("error: %v", err)
		// send bucket or retry process
	}

	// send message to topic FinishRetryProcessTopic, to finish retry process.
}

func buildUrl(host string, path string) string {
	return fmt.Sprintf("%s%s", host, path)
}
