package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gurbaj5124871/realtime-stream-processor/db"
	"github.com/gurbaj5124871/realtime-stream-processor/kafka"
	"github.com/gurbaj5124871/realtime-stream-processor/temporal"
)

func main() {
	ctx := context.Background()

	// Connect to database
	mongoURI := "mongodb://localhost:27017"
	mongoDBName := "sample-db"
	db.InitialiseMongo(ctx, &mongoURI, &mongoDBName)

	// Create temporal client and worker
	tc := temporal.InitialiseTemporalClient()
	go temporal.InitialiseTemporalWorker(tc)

	// Initialise and start kafka consumer
	kafkaTopic := "sample-topic"
	kc := kafka.GetKafkaConsumer(&kafkaTopic)
	go kafka.ConsumeDataStream(ctx, kc, tc)

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down after cleanup")

	// Close kafka consumers
	err := kc.Close()
	if err != nil {
		log.Printf("Error while closing kafka consumer: %s\n", err)
	}

	// Close temporal worker and client
	(*tc).Close()

	// Disconnect from database
	db.DisconnectMongo(ctx)
}
