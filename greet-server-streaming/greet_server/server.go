package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/y00rb/grpc-hands-on/greet-server-streaming/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Greet: Say greet for user
func (*server) GreetManyTimes(req *greetpb.GreetRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for index := 0; index < 10; index++ {
		result := "Hello " + firstName + " This is number " + strconv.Itoa(index) + "message on one request"
		res := &greetpb.GreetResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
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
