package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Yeremi528/grpc-golang/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("Do prime was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading streaming: %v", err)
		}

		log.Printf("Primes: %d", res.Result)
	}

}
