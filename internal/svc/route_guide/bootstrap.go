package routeguide

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/captain-corgi/grpc-go-example/api/routeguide"
	"github.com/captain-corgi/grpc-go-example/internal/svc/route_guide/usecase"
)

//RegisterGRPCService register service with gRPC server
func RegisterGRPCService(sv *grpc.Server) {
	// Register RouteGuide service
	pb.RegisterRouteGuideServer(sv, usecase.NewRouteGuideServer())
}

//RegisterGateway register services with gRPC gateway
func RegisterGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	// Register RouteGuide service
	//err := pb.RegisterRouteGuideHandlerFromEndpoint(ctx, mux, endpoint, opts)
	// if err != nil {
	// 	log.Fatalf("Failed to register gateway: %v", err)
	// }
}
