package main

import (
	"log"
	"net"

	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	tags "github.com/babaunba/project-management/api-gateway/gen/tags/v1"
	"github.com/babaunba/project-management/api-gateway/internal/domain"
	"github.com/babaunba/project-management/api-gateway/internal/server"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	srv := grpc.NewServer()
	{
		s, err := server.New(domain.New(), client.Options{})
		if err != nil {
			log.Fatalf("failed to create server: %v", err)
		}
		tags.RegisterTagsServer(srv, s)
	}

	reflection.Register(srv)

	lis, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Fatalf("failed to serve: %v", srv.Serve(lis))
}
