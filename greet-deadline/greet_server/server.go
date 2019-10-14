package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/y00rb/grpc-hands-on/greet-deadline/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

// Greet: Say greet for user
func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	for index := 0; index < 3; index++ {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("The deadline exceeded!")
			return nil, status.Error(codes.DeadlineExceeded, "The deadline exceeded")
		}
		if ctx.Err() == context.Canceled {
			fmt.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello, " + firstName

	reps := &greetpb.GreetResponse{
		Result: result,
	}
	return reps, nil
}

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
