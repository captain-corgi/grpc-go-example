package main

import (
	"flag"

	"github.com/captain-corgi/grpc-go-example/internal/svc"
)

var (
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	// Http port
	rpcPort = ":8180"
)

func main() {
	flag.Parse()

	go func() {
		// Start grpc service
		svc.StartServices(*grpcServerEndpoint)
	}()
	go func() {
		// Start grpc gateway, support restful API
		svc.StartGateway(rpcPort, *grpcServerEndpoint)
	}()
	select {}
}
