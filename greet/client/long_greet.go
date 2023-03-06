package main

import (
	"context"
	"log"
	"strconv"
	"time"

	pb "joyful.go/go-grpc/greet/proto"
)

func doLongGreet(con pb.GreetServiceClient) {

	log.Printf("the doLongGreet has been invoked %v\n", con)

	stream, err := con.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error occurred during client stream operation with the following details :\n %v \n", err)
	}

	for i := 0; i < 10; i++ {
		req := &pb.GreetRequest{
			FirstName: "Joyful round " + strconv.Itoa(i),
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error occurred during client stream sending to server with the following details: \n %v \n", err)
		}

		time.Sleep(1 * time.Second)

	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error occurred during receive operation with the following information :\n%v\n", err)
	}

	log.Printf("The result receive from server is \n%s\n", response.Result)

}
