package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "joyful.go/go-grpc/greet/proto"
)

func doGreet_with_deadline(client pb.GreetServiceClient, timeout time.Duration) {

	log.Println("invoked doGreet_with_deadline")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Joyful",
	}

	res, err := client.GreetWithDeadLine(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceed, %v", e)
				return
			} else {
				log.Fatalf("gRPC server returned error, %v\n", err)
			}

		} else {

			log.Fatalf("Error occurred during initiating connection with gRPC server, \n%v\n", err)
		}

	}

	log.Printf("The result is %s", res.Result)

}
