package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	pb "joyful.go/go-grpc/blog/proto"
)

var addr string = "0.0.0.0:50051"
var collection *mongo.Collection

func main() {
	var connectionString string = "mongodb://localhost:27017/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000"

	// clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")
	clientOptions := options.Client().ApplyURI(connectionString)
	// client, err := mongo.Connect(context.Background(), clientOptions)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error occurred during establishing connection to mongo db server, %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// database := client.Database("blogdb")
	// collection = database.Database().Collection("blog")
	collection = client.Database("blogdb-go").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Fail to listen to address %v\n", err)
	}

	log.Printf("Listening on address %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Fail to serve : %v\n", err)
	}

}

type Server struct {
	pb.BlogServiceServer
}
