package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(reader *kafka.Reader) *KafkaConsumer {
	return &KafkaConsumer{reader: reader}
}

func (kc *KafkaConsumer) Consume() string {
	ctx := context.Background()
	msg, err := kc.reader.ReadMessage(ctx)
	if err != nil {
		log.Fatalf("Error while reading message: %v", err)
	}
	return string(msg.Value)
}
