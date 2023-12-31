package main

import (
	"log"
	"net"

	pb "github.com/Yeremi528/grpc-golang/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
