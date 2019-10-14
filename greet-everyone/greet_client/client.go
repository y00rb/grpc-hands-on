package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/y00rb/grpc-hands-on/greet-everyone/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %w", err)
	}

	c := greetpb.NewGreetServiceClient(conn)

	respStream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
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

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			fmt.Printf("sending req: %v\n", req)
			respStream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		respStream.CloseSend()
	}()

	go func() {
		for {
			resp, err := respStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error when receiving: %v", err)
				break
			}
			fmt.Printf("Receiving final Response: %v \n", resp.GetResult())
		}

		close(waitc)
	}()
	<-waitc

}
