package temporal

import (
	"log"

	"go.temporal.io/sdk/client"
)

func InitialiseTemporalClient() *client.Client {
	tc, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort, // localhost:7233
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client", err)
	}

	log.Print("Temporal Client created")

	return &tc
}
