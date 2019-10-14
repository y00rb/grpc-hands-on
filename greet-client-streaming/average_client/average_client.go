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

	respStream, err := c.Average(context.Background())
	if err != nil {
		log.Fatal((""))
	}

	reqs := []*greetpb.CalculateRequest{
		&greetpb.CalculateRequest{
			Calculate: &greetpb.Calculate{
				Number: 3,
			},
		},
		&greetpb.CalculateRequest{
			Calculate: &greetpb.Calculate{
				Number: 5,
			},
		},
		&greetpb.CalculateRequest{
			Calculate: &greetpb.Calculate{
				Number: 9,
			},
		},
		&greetpb.CalculateRequest{
			Calculate: &greetpb.Calculate{
				Number: 54,
			},
		},
		&greetpb.CalculateRequest{
			Calculate: &greetpb.Calculate{
				Number: 23,
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
	fmt.Printf("LongGreet Response: %v \n", resp.Result)

}
