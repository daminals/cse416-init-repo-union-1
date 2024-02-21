package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode" // Replace "your-package-path" with the actual package path
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedMarketServiceServer
}

type FileRequest struct {
	ip   string
	port int
}

var fileRequestList = []FileRequest{
	{"127.0.0.1", 50052},
}

func (s *server) AddFileRequest(ctx context.Context, in *pb.FileHash) (*emptypb.Empty, error) {
	peerCtx, _ := peer.FromContext(ctx)
	log.Printf("Received: %v from consumer at %v", in.GetHash(), peerCtx.Addr.String())
	return &emptypb.Empty{}, nil
}

func (s *server) GetFileRequests(ctx context.Context, in *pb.FileHash) (*pb.FileRequestList, error) {
	log.Printf("Received: %v", in.GetHash())
	var requests []*pb.FileRequest
	for _, req := range fileRequestList {
		requests = append(requests, &pb.FileRequest{
			Ip:   req.ip,
			Port: int32(req.port),
		})
	}
	return &pb.FileRequestList{Requests: requests}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMarketServiceServer(s, &server{})
	log.Println("Market Server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// protoc --go_out=. --go-grpc_out=. example.proto
