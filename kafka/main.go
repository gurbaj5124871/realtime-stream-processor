package kafka

import (
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	LastOffset  int64 = -1 // The most recent offset available for a partition.
	FirstOffset int64 = -2 // The least recent offset available for a partition.
)

func GetKafkaConsumer(topic *string) *kafka.Reader {
	brokers := []string{"localhost:9092"}
	groupID := "consumer-group-1"

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	consumerConf := kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       *topic,
		StartOffset: LastOffset,
		Dialer:      dialer,
		MinBytes:    1, // same value of Shopify/sarama
		MaxBytes:    57671680,
	}

	consumer := kafka.NewReader(consumerConf)

	log.Printf("Kafka consumer initialised for topic: %s\n", *topic)

	return consumer
}
