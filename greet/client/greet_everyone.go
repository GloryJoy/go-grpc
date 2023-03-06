package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "joyful.go/go-grpc/greet/proto"
)

func doGreet_Everyone(connection pb.GreetServiceClient) {

	log.Println("doGreet_Everyone is invoked")

	arryReq := []string{
		"Joy",
		"John",
		"Bod",
		"Nick",
	}

	stream, err := connection.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error occurred during establishing stream connection from client \n%v\n", err)
	}

	waitc := make(chan struct{})

	go func() {

		for _, name := range arryReq {

			err = stream.Send(&pb.GreetRequest{
				FirstName: name,
			})

			if err != nil {
				log.Fatalf("Error occurred during sending name to gRPC server \n%v\n", err)
			}

			log.Printf("sending name %s", name)
			time.Sleep(1 * time.Second)

		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, resError := stream.Recv()

			if resError == io.EOF {
				break
			}

			if resError != nil {
				log.Printf("Error occurred during receiving from server \n%v\n", resError)
				break
			}

			log.Println(res.Result)

		}
		close(waitc)
	}()

	<-waitc

}
