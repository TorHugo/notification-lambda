package main

import (
	"github.com/joho/godotenv"
	"log"
	"notification-api/config"
	"notification-api/internal/adapters/http"
	messaging "notification-api/internal/adapters/messaging"
	"notification-api/internal/core/service"
	"os"
)

func main() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatalf("Error loading .env file: %v", errLoad)
	}
	var topic = os.Getenv("KAFKA_NOTIFICATION_TOPIC")

	consumer := config.NewKafkaReader(topic)
	kafkaConsumer := messaging.NewKafkaConsumer(consumer)
	httpClient := http.NewHttpClient()
	notificationService := service.NewNotificationService(kafkaConsumer, httpClient)
	notificationService.ProcessNotifications()
}
