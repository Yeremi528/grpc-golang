package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Yeremi528/grpc-golang/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("Do AVG was invoked in client")

	reqs := []*pb.PrimeRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg_Api(context.Background())

	if err != nil {
		log.Fatalf("Error while calling DoAvg with: %v", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from doAvg: %v", err)
	}

	log.Printf("doAvg: %s", res.Result)

}
