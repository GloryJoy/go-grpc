package main

import (
	"io"
	"log"

	pb "joyful.go/go-grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Average function on gRPC server has been invoked with the following server information \n%v\n", stream)

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Something went wrong while doing server side stream processing with the following error information \n%v\n", err)
		}

		log.Printf("receiving %d on server", req.Number)

		sum += int(req.Number)
		count++

	}
}
