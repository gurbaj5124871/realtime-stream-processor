package temporal

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func LiveStreamProcessingWorkflow(ctx workflow.Context, eventStr string) error {
	r := temporal.RetryPolicy{
		MaximumAttempts:    3,
		InitialInterval:    1 * time.Second,
		BackoffCoefficient: 2,
	}
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy:         &r,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Temporal workflow", "eventStr", eventStr)

	var result string
	err := workflow.ExecuteActivity(ctx, ProcessDataActivity, eventStr).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}

	logger.Info("workflow completed.")

	return nil
}
