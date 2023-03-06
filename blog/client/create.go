package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "joyful.go/go-grpc/blog/proto"
)

func doCreateBlog(c pb.BlogServiceClient) {

	log.Println("doCreateBlog is invoked")
	req := &pb.Blog{
		AuthorId: "Joy",
		Title:    "First Blog",
		Content:  "Hello World!",
	}

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.Internal {
				log.Printf("gRPC error %s\n", e.Message())
			} else {
				log.Printf("Other gRPC error %s\n", e.Err())
			}

		} else {

			log.Fatalf("Error with the following information %v\n", err)
		}

		log.Printf("The Blog is created with the following object id : %s", res.Id)
	}

}
