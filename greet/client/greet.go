package main

import (
	"context"
	"log"

	pb "joyful.go/go-grpc/greet/proto"
)

func doGreet(con pb.GreetServiceClient) {
	log.Printf("doGreet has been invoked with the following input \n %v \n", con)
	response, err := con.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Joyful",
	})

	if err != nil {
		log.Fatalf("Error occured during grpc call at doGreet \n %v \n", err)
	}
	log.Printf("The call result is \n %s \n", response.Result)

}
