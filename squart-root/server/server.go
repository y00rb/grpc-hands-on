package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/y00rb/grpc-hands-on/squart-root/squartpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquartRoot(ctx context.Context, req *squartpb.SquartRootRequest) (*squartpb.SquartRootResponse, error) {
	fmt.Println("Received squart RPC")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &squartpb.SquartRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal("Failed to lisent: %w", err)
	}

	s := grpc.NewServer()
	squartpb.RegisterSquartRootServer(s, &server{})
	if serverErr := s.Serve(lis); serverErr != nil {
		log.Fatal("Failed to listen: %w", err)
	}
}
