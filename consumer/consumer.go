package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	port = flag.Int("port", 50052, "The consumer service port")
)

type server struct {
	pb.UnimplementedConsumerServiceServer
}

func (s *server) ReceiveFileInfo(ctx context.Context, in *pb.FileLink) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in)
	return nil, nil
}


func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMarketServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.AddFileRequest(ctx, &pb.FileHash{Hash: "hash"})
	if err != nil {
		log.Fatalf("could not add file request: %v", err)
	}
	log.Printf("Made file request!")

	// now i open a grpc connection for the producer to reach out to me
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterConsumerServiceServer(s, &server{})
	log.Printf("Consumer Server listening on port %d...\n", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
