package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/captain-corgi/grpc-go-example/api/proto/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type (
	//server to implement helloworld.GreeterServer
	server struct {
		helloworld.UnimplementedGreeterServer
	}
)

func (r *server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (reply *helloworld.HelloReply, err error) {
	log.Printf("Received: %v", request.GetName())
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello %s", request.GetName()),
	}, nil
}

func main() {
	// Define server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sv := grpc.NewServer()

	// Register service
	helloworld.RegisterGreeterServer(sv, &server{})

	// Start server
	log.Printf("Server started in port %s\n", port)
	log.Fatalf("failed to start server %v", sv.Serve(lis))
}
