package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "joyful.go/go-grpc/blog/proto"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {

	log.Printf("CreateBlog is invoked with \n%v\n", in)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error %v\n", err))
	}

	objectId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Failed to convert inserted ObjectID")
	}

	return &pb.BlogId{
		Id: objectId.Hex(),
	}, nil

}
