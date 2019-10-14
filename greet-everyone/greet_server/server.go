package main

import (
	"io"
	"log"
	"net"

	"github.com/y00rb/grpc-hands-on/greet-everyone/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Greet: Say greet for user
func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + " ! "
		sendErr := stream.Send(&greetpb.GreetResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
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
