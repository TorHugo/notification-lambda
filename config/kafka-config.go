package messaging

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Consumer() (*kafka.Consumer, error) {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatalf("Error loading .env file: %v", errLoad)
	}

	var host = os.Getenv("KAFKA_HOST")
	var group = os.Getenv("KAFKA_GROUP")
	port, err := strconv.Atoi(os.Getenv("KAFKA_PORT"))
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fullHost(host, port),
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	return consumer, nil
}

func fullHost(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
