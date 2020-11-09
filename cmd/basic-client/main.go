package main

import (
	"context"
	"log"
	"time"

	"github.com/captain-corgi/grpc-go-example/api/routeguide"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8180"
)

func main() {
	// Setup a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Register a client with server
	basicClient := routeguide.NewRouteGuideClient(conn)

	// Call gRPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := basicClient.GetFeature(ctx, &routeguide.Point{Latitude: 1, Longitude: 1})
	if err != nil {
		log.Fatalf("could not get feature: %v", err)
	}
	log.Printf("Feature: %+v", r)
}
