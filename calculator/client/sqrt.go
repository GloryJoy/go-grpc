package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "joyful.go/go-grpc/calculator/proto"
)

func doSqrt(client pb.CalculatorServiceClient) {

	log.Println("doSqrt is invoked")

	req := &pb.SqrtRequest{
		Number: -20,
	}

	res, err := client.Sqrt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("gRPC error occurred on gRPC server with code %s\n %s", e.Code(), e.Message())
			if e.Code() == codes.InvalidArgument {
				log.Printf("InvalidArgument %d, this number is invalid.", req.Number)
				return
			}

		} else {

			log.Fatalf("None gRPC error occurred in gRpc server call \n%v\n", err)
		}
	}

	log.Printf("The root of the given number %d is %d", req.Number, res.Result)

}
