package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/y00rb/grpc-hands-on/greet-client-streaming/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %w", err)
	}

	c := greetpb.NewGreetServiceClient(conn)

	respStream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal((""))
	}

	reqs := []*greetpb.GreetRequest{
		&greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Paul",
			},
		},
		&greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Zac",
			},
		},
		&greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Gary",
			},
		},
	}

	for _, req := range reqs {
		fmt.Printf("sending req: %v\n", req)
		respStream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	resp, err := respStream.CloseAndRecv()
	if err != nil {
		log.Fatal("")
	}
	fmt.Printf("LongGreet Response: %v \n", resp)

}
