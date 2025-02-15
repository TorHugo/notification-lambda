package config

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
)

func NewKafkaReader(topic string) *kafka.Reader {
	var host = os.Getenv("KAFKA_HOST")
	var groupID = os.Getenv("KAFKA_GROUP")
	port, err := strconv.Atoi(os.Getenv("KAFKA_PORT"))
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	log.Printf("Connecting to Kafka Brokers, host: %s, group: %s, port: %d\n", host, groupID, port)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{fullHost(host, port)},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func fullHost(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
