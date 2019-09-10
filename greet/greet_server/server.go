package main

import (
	"context"
	"log"
	"net"

	"github.com/y00rb/grpc-hands-on/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Greet: Say greet for user
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello, " + firstName

	reps := &greetpb.GreetResponse{
		Result: result,
	}
	return reps, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen: %w", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to listen: %w", err)
	}
}
