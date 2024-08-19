package messaging

import (
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(c *kafka.Consumer) *KafkaConsumer {
	return &KafkaConsumer{
		consumer: c,
	}
}

func (kc *KafkaConsumer) Consume(topic string) string {

	err := kc.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to Kafka topic: %v", err)
	}

	timeout := 10 * time.Second
	ev := kc.consumer.Poll(int(timeout.Milliseconds()))
	if ev == nil {
		log.Fatalf("Timeout when waiting for message from topic")
	}

	switch e := ev.(type) {
	case *kafka.Message:
		return string(e.Value)
	case kafka.Error:
		log.Fatalf("Kafka error: %v", e)
	default:
		log.Fatalf("Unexpected event type from Kafka: %v", e)
	}
	return ""
}
