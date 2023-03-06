package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "joyful.go/go-grpc/calculator/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Fail to listen to address %v\n", err)
	}

	log.Printf("Listening on address %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Fail to serve : %v\n", err)
	}

}

type Server struct {
	pb.CalculatorServiceServer
}
