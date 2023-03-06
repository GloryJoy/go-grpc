package main

import (
	"context"
	"log"

	pb "joyful.go/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function is invoked with the following input parameter: \n %+v \n", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil

}
