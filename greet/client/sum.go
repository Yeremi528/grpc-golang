package main

import (
	"context"
	"log"

	pb "github.com/Yeremi528/grpc-golang/greet/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Printf("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 50,
		B: 30,
	})

	if err != nil {
		log.Fatalf("Could not Sum: %v", err)
	}

	log.Printf("Sum: %v", res.C)
}
