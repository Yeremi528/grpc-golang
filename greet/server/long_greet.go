package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Yeremi528/grpc-golang/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		res += fmt.Sprintf("Hello %s", req.FirstName)
	}

}
