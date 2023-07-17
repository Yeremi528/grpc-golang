package main

import (
	"context"
	"log"

	pb "github.com/Yeremi528/grpc-golang/greet/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with; %v", in)

	return &pb.SumResponse{
		C: in.A + in.B,
	}, nil
}
