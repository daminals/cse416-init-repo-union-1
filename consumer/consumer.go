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
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type server struct {
	pb.UnimplementedConsumerServiceServer
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

	// parse the port from the addr string
	// _, port, err := net.SplitHostPort(*addr)
	// if err != nil {
	// 	log.Fatalf("Failed to parse port: %v", err)
	// }
	// // convert the port string to an int
	// portInt, err := net.LookupPort("tcp", port)

	// if err != nil {
	// 	log.Fatalf("Failed to convert port to int: %v", err)
	// }
	portInt := 50052

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.AddFileRequest(ctx, &pb.FileHash{Hash: "hash"})
	if err != nil {
		log.Fatalf("could not add file request: %v", err)
	}
	log.Printf("Made file request!")

	// now i open a grpc connection for the producer to reach out to me
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", portInt))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterConsumerServiceServer(s, &server{})
	log.Printf("Market Server listening on port %d...\n", portInt)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
