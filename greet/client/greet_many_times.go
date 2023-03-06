package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "joyful.go/go-grpc/greet/proto"
)

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes has been invoked")

	requestMessage := &pb.GreetRequest{
		FirstName: "Joyful",
	}

	stream, err := client.GreetManyTimes(context.Background(), requestMessage)

	if err != nil {
		log.Fatalf("Error occurred connecting to stream Greet Many Times %v\n", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("Breaking out of the loop because the stream is ended.")
			break
		}

		if err != nil {
			log.Fatalf("Fatal error occurred during reading the stream with the following information \n %v \n", err)
		}

		fmt.Printf("Printing the result from reading the stream -- %s", response.Result)
	}
}
