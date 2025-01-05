package api

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"

	tags "github.com/babaunba/project-management/api-gateway/gen/tags/v1"
)

// there is no reason to pass req/resp as pointers on account of temporal deserialising the data
type domain interface {
	GenerateTagsWF(workflow.Context, *tags.GenerateTagsRequest) (*tags.GenerateTagsResponse, error)
}

// Server is a gRPC server implementation that requires workflow domain definition
type Server struct {
	tags.UnimplementedTagsServer
	domain domain
	client client.Client // temporal client
}

// New is a *Server constructor
func New(domain domain, opts client.Options) (srv *Server, err error) {
	client, err := client.Dial(opts)
	srv = &Server{domain: domain, client: client}
	return
}
