package main

import (
	"fmt"
	"io"
	"log"
	"strconv"

	pb "joyful.go/go-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("Long greet server service has been invoked with the following value %v\n", stream)

	responseMsg := ""
	var count int16 = 0
	for {

		msgReq, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: responseMsg,
			})
		}

		if err != nil {
			log.Fatalf("Fatal error occured during server stream operation with the following information %v\n", err)
		}

		log.Printf("Printing...%s\n", strconv.Itoa(int(count)))
		count++

		log.Printf("The receiving value is \n%s\n", responseMsg)
		responseMsg += fmt.Sprintf("Hello %s\n", msgReq.FirstName)

	}
	return nil

}
