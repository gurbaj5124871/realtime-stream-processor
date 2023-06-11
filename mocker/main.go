package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/segmentio/kafka-go"
)

type EventDataUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type EventData struct {
	User EventDataUser
}

type Event struct {
	EventType    string
	EventVersion string
	Data         EventData
}

func main() {

	brokers := []string{"localhost:9092"}

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	consumerConf := kafka.WriterConfig{
		Brokers: brokers,
		Dialer:  dialer,
	}

	producer := kafka.NewWriter(consumerConf)
	msgs := []kafka.Message{}
	topic := "sample-topic"

	for i := 0; i < 1000; i += 1 {
		event := Event{
			EventType:    "sign-up",
			EventVersion: "1.0",
			Data: EventData{
				User: EventDataUser{
					ID:        int64(i + 1),
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
				},
			},
		}
		json, err := json.Marshal(event)
		if err != nil {
			fmt.Printf("Error while marshalling event: %s\n", err)
			panic(err)
		}

		msg := kafka.Message{
			Key:   []byte("Key-" + fmt.Sprintf("%d", i+1)),
			Value: json,
			Topic: topic,
		}
		msgs = append(msgs, msg)
	}

	err := producer.WriteMessages(context.Background(), msgs...)
	if err != nil {
		fmt.Printf("Error while writing messages to kafka: %s\n", err)
	}
}
