package temporal

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/gurbaj5124871/realtime-stream-processor/db"
	"go.temporal.io/sdk/activity"
)

type LiveStreamEventDataUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LiveStreamEventData struct {
	User *LiveStreamEventDataUser `json:"user,omitempty"`
}

type LiveStreamEvent struct {
	EventType    string              `json:"eventType"`
	EventVersion string              `json:"eventVersion"`
	Data         LiveStreamEventData `json:"data"`
}

func ProcessDataActivity(ctx context.Context, eventStr string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "event", eventStr)

	// Logic to process/transform data
	// Filter out events with missing data
	event := LiveStreamEvent{}
	if err := json.Unmarshal([]byte(eventStr), &event); err != nil {
		logger.Error("Activity", "error", err)
		return err
	}
	if event.Data.User == nil {
		logger.Error("Activity", "error", "missing user data, user:: %+v", event.Data.User)
		return errors.New("missing user data")
	}

	// Write event into database
	dbEvent := db.LiveStreamEvent{
		ID:        event.Data.User.ID,
		FirstName: event.Data.User.FirstName,
		LastName:  event.Data.User.LastName,
		Timestamp: time.Now(),
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := dbEvent.InsertOne(ctx); err != nil {
		logger.Error("Activity", "error", err)
		return err
	}

	return nil
}
