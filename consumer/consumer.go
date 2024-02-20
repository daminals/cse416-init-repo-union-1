package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	port = flag.Int("port", 50052, "The consumer service port")
	fileResponse = &fileInfo{}
	srv *grpc.Server
)

type server struct {
	pb.UnimplementedConsumerServiceServer // this is the consumer service
}

type fileInfo struct {
	Link           string // url to download the file
	Token          string // access token to download the file
	PaymentAddress string // payment address to send the payment
}

// implement a function in fileinfo to check if the file info is complete
func (f *fileInfo) isComplete() bool {
	return f.Link != "" && f.Token != "" && f.PaymentAddress != ""
}

// RecieveFileInfo is the function that the producer will call to send the file info
// afterwards, the consumer should close the server and make an http request to the producer
// to download the file
func (s *server) ReceiveFileInfo(ctx context.Context, in *pb.FileLink) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in)

	fileResponse.Link = in.GetLink()
	fileResponse.Token = in.GetToken()
	fileResponse.PaymentAddress = in.GetPaymentAddress()

	// Close the server
	srv.Stop()

	// For now, just return an empty response
	return &emptypb.Empty{}, nil
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

	srv = grpc.NewServer()
	pb.RegisterConsumerServiceServer(srv, &server{})
	log.Printf("Consumer Server listening on port %d...\n", *port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	// check if the file info is complete
	if !fileResponse.isComplete() {
		// fatal crash if any of the file info is missing
		log.Fatalf("File info is missing")
	}

	// send an http request to the producer to download the file
	netClient := &http.Client{}
	req, err := http.NewRequest("GET", fileResponse.Link, nil)
	if err != nil {
		log.Fatalf("Error creating http request: %v", err)
	}
	req.Header.Set("Authorization", fileResponse.Token)
	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending http request: %v", err)
	}
	// check if the response is 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error downloading file: %v", resp.Status)
	}
}
