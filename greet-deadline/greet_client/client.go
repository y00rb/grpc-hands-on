package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/y00rb/grpc-hands-on/greet-deadline/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %w", err)
	}

	c := greetpb.NewGreetServiceClient(conn)
	res := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "y00rb",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	resp, err := c.GreetWithDeadline(ctx, res)

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: %v", statusErr)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)

		}
		return
	}
	fmt.Printf("Response from Greet: %v", resp.Result)
}
