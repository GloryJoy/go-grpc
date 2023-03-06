package main

import (
	"log"

	pb "joyful.go/go-grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	log.Printf("Prime function is invoked with the following data %v\n", in)

	number := in.Number
	diviser := int64(2)

	for number > 1 {
		if number%diviser == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: diviser,
			})

			number /= diviser
		} else {
			diviser++
		}
	}

	return nil
}
