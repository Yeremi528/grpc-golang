package main

import (
	"io"
	"log"
	"sort"

	pb "github.com/Yeremi528/grpc-golang/calculator/proto"
)

func (s *Server) Max_Api(stream pb.CalculatorService_Max_ApiServer) error {
	log.Printf("Max_Api was invoked")

	var mySlice []int64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error in Max_API while reading client: %v", err)
		}

		if len(mySlice) != 0 {
			if mySlice[0] < req.Number {
				err = stream.Send(&pb.PrimeResponse{Result: int64(req.Number)})

				if err != nil {
					log.Fatalf("Error while sending data to client: %v", err)
				}
			}
		}

		mySlice = append(mySlice, req.Number)

		if len(mySlice) == 1 {
			err = stream.Send(&pb.PrimeResponse{Result: int64(req.Number)})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
			}
		}

		sort.Slice(mySlice, func(i, j int) bool { return mySlice[i] > mySlice[j] })

	}

}
