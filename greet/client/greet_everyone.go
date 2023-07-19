package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Yeremi528/grpc-golang/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	waitG := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v time:%v", req, time.Now().Nanosecond())

			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v", err)
			}

			log.Printf("Received: %v time:%v", res.Result, time.Now().Nanosecond())
		}
		close(waitG)

	}()

	<-waitG

}
