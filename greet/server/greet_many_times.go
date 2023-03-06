package main

import (
	"fmt"
	"log"

	pb "joyful.go/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	log.Printf("Greet Many Times has been invoked \n %v \n", in)

	for i := 0; i < 10; i++ {
		responseResult := fmt.Sprintf("Hello, %s and the number of time is %d \n", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: responseResult,
		})
	}

	return nil

}
