package main

import (
	"context"
	"log"

	pb "github.com/Yeremi528/grpc-golang/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})

	if err != nil {
		log.Fatalf("Could no greet: %v", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
