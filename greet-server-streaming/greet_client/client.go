package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/y00rb/grpc-hands-on/greet-server-streaming/greetpb"
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
	respStream, err := c.GreetManyTimes(context.Background(), res)
	if err != nil {
		log.Fatal("error while calling Green: %w", err)
	}
	for {
		msg, err := respStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error while calling Green: %w", err)
		}
		fmt.Printf("Response from Greet: %v\n", msg.GetResult())
	}

}
