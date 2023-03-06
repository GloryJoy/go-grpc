package main

import (
	"io"
	"log"

	pb "joyful.go/go-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("invoking Max function on gRPC server")

	var max int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error occurred on gRPC Server, Max function --> \n%v\n", err)
		}

		log.Printf("receiving %d", req.Number)

		if req.Number > max {
			max = req.Number
			log.Printf("the new max is %d", max)
		}

		res := &pb.MaxResponse{
			Result: max,
		}
		err = stream.Send(res)
		if err != nil {
			log.Fatalf("Error occurred during returning max to client \n%v\n", err)
		}

	}

}
