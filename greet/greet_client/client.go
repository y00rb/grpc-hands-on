package main

import (
	"context"
	"fmt"
	"log"

	"github.com/y00rb/grpc-hands-on/greet/greetpb"
	"google.golang.org/grpc"
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
	resp, err := c.Greet(context.Background(), res)
	if err != nil {
		log.Fatal("error while calling Green: %w", err)
	}
	fmt.Printf("Response from Greet: %v", resp.Result)
}
