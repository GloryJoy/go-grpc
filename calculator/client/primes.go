package main

import (
	"context"
	"io"
	"log"

	pb "joyful.go/go-grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("The doPrimes calculator service client has been invoked with the following value %v\n", c)

	messageRequest := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), messageRequest)

	if err != nil {
		log.Fatalf("Fatal error in the Primes stream client as the following details \n%s\n", err.Error())
	}

	for {

		msgResponse, err := stream.Recv()

		if err == io.EOF {
			break
		}

		log.Printf("The result from server in stream mode is : %d", msgResponse.Result)
	}

}
