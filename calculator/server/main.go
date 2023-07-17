package main

import (
	"log"
	"net"

	pb "github.com/Yeremi528/grpc-golang/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf("Server failed with: %v", err)
	}

	log.Printf("Server running in port: %v", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
