package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Yeremi528/grpc-golang/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {

	reqs := []pb.PrimeRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	stream, err := c.Max_Api(context.Background())

	if err != nil {
		log.Fatalf("Errow while invoking stream: %v", err)
	}

	func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v", req.Number)
			stream.Send(&pb.PrimeRequest{Number: req.Number})
		}
		stream.CloseSend()
	}()
	func() {
		for {
			req, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("err reading req: %v", err)
			}

			log.Printf("we reicibing: %v", req.Result)
		}
	}()

}
