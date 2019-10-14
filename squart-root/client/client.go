package main

import (
	"context"
	"fmt"
	"log"

	"github.com/y00rb/grpc-hands-on/squart-root/squartpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, connErr := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if connErr != nil {
		log.Fatal("could not connect: %w", connErr)
	}

	c := squartpb.NewSquartRootClient(conn)
	res := &squartpb.SquartRootRequest{
		Number: 5,
	}
	resp, err := c.SquartRoot(context.Background(), res)
	fmt.Println(err != nil)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println(("We probably send a negative number!"))
				return
			}
		} else {
			log.Fatalf("Big error calling squart %v", err)
			return
		}
	}
	fmt.Printf("Response from square root: %v\n", resp.GetNumberRoot())
}
