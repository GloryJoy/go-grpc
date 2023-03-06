package main

import (
	"context"
	"log"

	pb "joyful.go/go-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("The Greet function has been invoked with the following input, %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello --> " + in.FirstName,
	}, nil
}
