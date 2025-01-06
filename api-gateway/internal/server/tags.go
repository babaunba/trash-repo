package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tags "github.com/babaunba/project-management/api-gateway/gen/tags/v1"
)

const (
	startWFID = "tags-workflow"
	taskQueue = "tags-tasks"
)

// GenerateTags adds a tags generation task to the queue and expects a worker to complete it
func (s *Server) GenerateTags(ctx context.Context, req *tags.GenerateTagsRequest) (resp *tags.GenerateTagsResponse, err error) {
	opts := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("%s:%s", startWFID, uuid.New().String()),
		TaskQueue: taskQueue,
	}

	we, err := s.client.ExecuteWorkflow(ctx, opts, s.domain.GenerateTagsWF, req)
	if err != nil {
		err = fmt.Errorf("unable to execute workflow: %w", err) // first wrap and add a message
		err = status.Error(codes.Internal, err.Error())         // then use a gRPC status wrapper
		return
	}

	err = we.Get(ctx, &resp)
	if err != nil {
		err = fmt.Errorf("unable to get workflow result: %w", err) // analogous
		err = status.Error(codes.Internal, err.Error())
		return
	}

	return
}
