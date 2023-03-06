package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "joyful.go/go-grpc/greet/proto"
)

func (s *Server) GreetWithDeadLine(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello, " + in.FirstName,
	}, nil

}
