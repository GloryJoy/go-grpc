package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "joyful.go/go-grpc/greet/proto"
)

var addr string = "localhost:50051"

func main() {

	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server address : %s, with the following error \n %v\n", addr, err)
	}

	defer connection.Close()
	log.Printf("The client connected to grpc server.")

	clientApp := pb.NewGreetServiceClient(connection)

	// doGreet(clientApp)
	// doGreetManyTimes(clientApp)
	// doLongGreet(clientApp)
	// doGreet_Everyone(clientApp)
	doGreet_with_deadline(clientApp, 2*time.Second)

}
