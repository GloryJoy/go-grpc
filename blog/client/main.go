package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "joyful.go/go-grpc/blog/proto"
)

var addr string = "localhost:50051"

func main() {

	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server address : %s, with the following error \n %v\n", addr, err)
	}

	defer connection.Close()
	log.Printf("The client connected to grpc server.")

	clientApp := pb.NewBlogServiceClient(connection)

	// doCalculation(clientApp)
	// doPrimes(clientApp)
	// doAvg(clientApp)

	// doMax(clientApp)
	// doSqrt(clientApp)
	doCreateBlog(clientApp)

}
