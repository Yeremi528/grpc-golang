package main

import (
	"io"
	"log"

	pb "github.com/Yeremi528/grpc-golang/calculator/proto"
)

func (s *Server) Avg_Api(stream pb.CalculatorService_Avg_ApiServer) error {
	log.Printf("Avg function was invoked")

	res := float32(0)

	var mySlice []int64
	for {

		req, err := stream.Recv()

		if err == io.EOF {
			for i := int64(0); i < int64(len(mySlice)); i++ {
				res += float32(mySlice[i])
			}
			return stream.SendAndClose(&pb.AvgResponse{
				Result: res / 4,
			})
		}

		if err != nil {
			log.Fatalf("err while calling avgApi with: %v", err)
		}

		mySlice = append(mySlice, req.Number)
	}

}
