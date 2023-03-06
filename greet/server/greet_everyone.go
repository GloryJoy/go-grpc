package main

import (
	"io"
	"log"

	pb "joyful.go/go-grpc/greet/proto"
)

func (s *Server) GreetEveryone(greetServer pb.GreetService_GreetEveryoneServer) error {

	log.Printf("Greet Everyone gRPC Server function is invoked")

	for {
		greetReq, err := greetServer.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error occurred on gRPC server side, \n%v\n", err)
		}

		log.Printf("Receiving %s", greetReq.FirstName)

		greetRes := &pb.GreetResponse{
			Result: "Nice to meet you, " + greetReq.FirstName,
		}

		err = greetServer.Send(greetRes)
		if err != nil {
			log.Fatalf("Error occurred during greet response stream \n%v\n", err)
		}

		log.Printf("Sending %s", greetRes.Result)

	}

}
