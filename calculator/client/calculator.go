package main

import (
	"context"
	"log"

	pb "joyful.go/go-grpc/calculator/proto"
)

func doCalculation(con pb.CalculatorServiceClient) {
	log.Printf("doCalculation has been invoked with the following input \n %v \n", con)
	response, err := con.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  100,
		SecondNumber: 100,
	})

	if err != nil {
		log.Fatalf("Error occured during grpc call at doCalculation \n %v \n", err)
	}
	log.Printf("The call result is \n %d \n", response.Result)

}
