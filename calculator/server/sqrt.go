package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "joyful.go/go-grpc/calculator/proto"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {

	log.Println("Sqrt function on gRPC server is invoked")

	number := in.Number
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number %d", in.Number),
		)
	}

	return &pb.SqrtResponse{
		Result: int64(math.Sqrt(float64(number))),
	}, nil

}
