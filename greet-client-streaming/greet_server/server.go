package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/y00rb/grpc-hands-on/greet-client-streaming/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.GreetResponse{
				Result: result,
			})
		}
		if err != nil {
			fmt.Printf("Error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + " ! "
	}
}

func (*server) Average(stream greetpb.GreetService_AverageServer) error {
	fmt.Printf("Average function was invoked with a streaming request")
	result := 0.
	streamLength := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.CalculateResponse{
				Result: float32(result / streamLength),
			})
		}
		if err != nil {
			fmt.Printf("Error while reading client stream: %v", err)
		}
		result += float64(req.GetCalculate().GetNumber())
		streamLength++
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
