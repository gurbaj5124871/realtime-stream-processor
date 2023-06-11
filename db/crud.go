package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type LiveStreamEvent struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Timestamp time.Time `json:"timestamp"`
}

func (l *LiveStreamEvent) InsertOne(ctx context.Context) error {
	bsonEvent, err := bson.Marshal(l)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	coll, err := DB.Collection("liveStreamEvents").InsertOne(ctx, bsonEvent)
	if err != nil {
		return err
	}
	if coll.InsertedID == nil {
		return errors.New("failed to insert event into database")
	}

	return nil
}
