package usecase

import (
	"context"
	"fmt"

	"github.com/captain-corgi/grpc-go-example/api/routeguide"
	pb "github.com/captain-corgi/grpc-go-example/api/routeguide"
)

//RouteGuideServer implement API RouteGuide
type RouteGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

//NewRouteGuideServer return new RouteGuide server
func NewRouteGuideServer() *RouteGuideServer {
	return &RouteGuideServer{}
}

//GetFeature return feature of selected point
func (r *RouteGuideServer) GetFeature(ctx context.Context, p *pb.Point) (*routeguide.Feature, error) {
	fmt.Printf("Received %+v\n", p)
	return &pb.Feature{
		Name:     "Dummy feature",
		Location: p,
	}, nil
}
