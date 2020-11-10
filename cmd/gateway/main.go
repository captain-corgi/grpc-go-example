package main

import (
	"flag"
	"github.com/captain-corgi/grpc-go-example/internal/svc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")

	// Http port
	rpcPort = ":8180"
)

func main() {
	flag.Parse()

	// Start grpc gateway, to support restful API
	svc.StartGateway(rpcPort, *grpcServerEndpoint)
}
