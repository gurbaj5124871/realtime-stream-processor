package temporal

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func LiveStreamProcessingWorkflow(ctx workflow.Context, name string) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Temporal workflow", "name", name)

	var result string
	err := workflow.ExecuteActivity(ctx, ProcessDataActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}

	logger.Info("workflow completed.")

	return nil
}
