package routeguide

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/captain-corgi/grpc-go-example/api/routeguide"
	"github.com/captain-corgi/grpc-go-example/internal/svc/route_guide/usecase"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8180, "Basic server port")
)

//Bootstrap start service
func Bootstrap() {
	// Define server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sv := grpc.NewServer()

	// Register service
	pb.RegisterRouteGuideServer(sv, usecase.NewRouteGuideServer())

	// Start server
	log.Printf("Server started in port %d\n", *port)
	log.Fatalf("failed to start server %v", sv.Serve(lis))
}
