package svc

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"

	routeGuideService "github.com/captain-corgi/grpc-go-example/internal/svc/route_guide"
	userService "github.com/captain-corgi/grpc-go-example/internal/svc/user"
)

//StartServices run all services
func StartServices(grpcEndpoint string) {
	// Define server
	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sv := grpc.NewServer()

	//Register services
	routeGuideService.RegisterGRPCService(sv)
	userService.RegisterGRPCService(sv)

	// Start server
	log.Printf("Server listening on in port %s\n", grpcEndpoint)
	log.Fatalf("failed to start server %v", sv.Serve(lis))
}

//StartGateway run all gRPC gateways
func StartGateway(rpcPort string, grpcEndpoint string) {
	defer glog.Flush()

	// Initial context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	userService.RegisterGateway(ctx, mux, grpcEndpoint, opts)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("Grpc server started on port %s\n", grpcEndpoint)
	log.Printf("Http server listening on port %s\n", rpcPort)
	glog.Fatal(http.ListenAndServe(rpcPort, mux))
}
