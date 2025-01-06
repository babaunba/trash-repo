package main

import (
	"log"

	w "github.com/babaunba/project-management/api-gateway/pkg/worker"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("failed to connect to temporal: %v", err)
	}

	w := w.RegisterTagsWorker(c, &worker.Options{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("failed to start worker: %v", err)
	}
}
