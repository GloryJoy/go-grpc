package main

import (
	"context"
	"log"

	pb "joyful.go/go-grpc/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {

	log.Printf("doAvg client side is invoked")

	numbers := []int64{
		10, 20, 30, 40,
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error occurred in client stream operation with following message: \n%v\n ", err)
	}

	for _, v := range numbers {
		log.Printf("sending %d to server", v)

		err := stream.Send(&pb.AvgRequest{Number: v})
		if err != nil {
			log.Fatalf("Error occurred during sending stream to server \n%v\n", err)
		}

	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error during closing and receiving stream from server \n%v\n", err)
	}

	log.Printf("The average result is %f", response.Result)

}
