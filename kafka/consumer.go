package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/gurbaj5124871/realtime-stream-processor/temporal"
	"github.com/segmentio/kafka-go"
	"go.temporal.io/sdk/client"
)

// The consumer will panic and crash on errors if we fail to read the message or fail to register into temporal
// Needs to implement a retry mechanism where consumers restarts on failure
func ConsumeDataStream(ctx context.Context, consumer *kafka.Reader, tc *client.Client) {
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		msg, err := consumer.ReadMessage(c)
		if err != nil {
			if err.Error() != "EOF" {
				log.Printf("Error while reading message from kafka: %s\n", err)
			}
			panic(err)
		}

		log.Printf("Message received: %s\n", string(msg.Value))
		workflowOptions := client.StartWorkflowOptions{
			ID:        msg.Topic + "-" + fmt.Sprintf("%d", msg.Partition) + "-" + fmt.Sprintf("%d", msg.Offset),
			TaskQueue: "livestream-data-processor-queue",
		}
		_, err = (*tc).ExecuteWorkflow(c, workflowOptions, temporal.LiveStreamProcessingWorkflow, string(msg.Value))
		if err != nil {
			panic(err)
		}
	}
}
