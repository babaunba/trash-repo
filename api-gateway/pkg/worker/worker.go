package worker

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/babaunba/project-management/api-gateway/internal/domain"
)

const (
	taskQueue = "tags-tasks"
)

func RegisterTagsWorker(c client.Client, opts *worker.Options) (w worker.Worker) {
	w = worker.New(c, taskQueue, *opts)
	w.RegisterWorkflow(domain.New().GenerateTagsWF)
	return
}
