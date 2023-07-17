package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Yeremi528/grpc-golang/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Yeremi",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the steam: %v", err)
		}

		log.Printf("GreetManyTimes: %s", msg.Result)
	}

}
