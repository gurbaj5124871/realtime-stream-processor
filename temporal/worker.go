package temporal

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func InitialiseTemporalWorker(temporalClient *client.Client) *worker.Worker {
	queue := "livestream-data-processor-queue"

	// use options to control max concurrent activity executions, retry policy and timeouts
	opts := worker.Options{}
	tw := worker.New(*temporalClient, queue, opts)
	defer tw.Stop()

	tw.RegisterWorkflow(LiveStreamProcessingWorkflow)
	tw.RegisterActivity(ProcessDataActivity)

	err := tw.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start temporal worker", err)
	}

	return &tw
}
