package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitialiseMongo(ctx context.Context, uri *string, dbName *string) *mongo.Database {
	mongoDb, err := mongo.NewClient(options.Client().ApplyURI(*uri))
	if err != nil {
		log.Fatalf("Error while connecting to mongodb: %s\n", err)
		os.Exit(1)
	}

	c, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = mongoDb.Connect(c)
	if err != nil {
		log.Fatalf("Error while connecting to mongodb: %s\n", err)
		os.Exit(1)
	}

	log.Println("Mongodb Connected")

	DB = mongoDb.Database(*dbName)
	return DB
}

func DisconnectMongo(ctx context.Context) {
	c, cancel := context.WithTimeout(ctx, 3)
	defer cancel()

	err := DB.Client().Disconnect(c)
	if err != nil {
		log.Printf("Error while disconnecting database: %s\n", err)
	}
}
