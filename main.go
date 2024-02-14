package main

import (
	"context"
	"log"
	"net"

	pb "github.com/daminals/cse416-init-repo-union-1/example" // Replace "your-package-path" with the actual package path
	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &server{})
	log.Println("Server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// protoc --go_out=. --go-grpc_out=. example.proto
