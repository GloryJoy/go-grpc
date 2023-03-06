package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "joyful.go/go-grpc/calculator/proto"
)

func doMax(client pb.CalculatorServiceClient) {
	log.Printf("Invoke doMax")

	waitc := make(chan struct{})

	arrReq := []int64{
		10, 20, 30, 40, 30, 50,
	}

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("Error occurred during establishing a call to Max function on the server \n%v\n", err)
	}

	go func() {

		for _, number := range arrReq {
			req := &pb.MaxRequest{
				Number: number,
			}

			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error occurred ruing sending request to gRPC Server \n%v\n", err)
			}
			log.Printf("Sending value is %d", req.Number)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error occurred during receiving stream from gRPC server \n%v\n", err)
				break
			}

			log.Printf("the receiving value is %d.", res.Result)

		}

		close(waitc)
	}()

	<-waitc

}
