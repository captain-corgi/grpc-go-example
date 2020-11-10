package user

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/captain-corgi/grpc-go-example/api/user"
	"github.com/captain-corgi/grpc-go-example/internal/svc/user/usecase"
)

//RegisterGRPCService register service with gRPC server
func RegisterGRPCService(sv *grpc.Server) {
	// Register user management service
	pb.RegisterUserServiceServer(sv, usecase.NewUserManagement())
}

//RegisterGateway register services with gRPC gateway
func RegisterGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	// Register User service
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}
}
