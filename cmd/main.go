package main

import (
	"log"
	config "notification-lambda/config"
	"notification-lambda/internal/adapters/http"
	"notification-lambda/internal/adapters/messaging"
	"notification-lambda/internal/core/service"
)

func main() {
	consumer, err := config.Consumer()
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	kafkaConsumer := messaging.NewKafkaConsumer(consumer)
	httpClient := http.NewHttpClient()
	notificationService := service.NewNotificationService(kafkaConsumer, httpClient)
	notificationService.ProcessNotifications()
}
