package internal

import (
	"context"
	"log"
	"time"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func MarketServerRequest(addr string) ([]*pb.FileRequest, error) {
	// Establish connection with the market server
	marketServerConnection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer marketServerConnection.Close()
	marketServerClient := pb.NewMarketServiceClient(marketServerConnection)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	marketServerResponse, err := marketServerClient.GetFileRequests(ctx, &pb.FileHash{Hash: "hash"})
	if err != nil {
		log.Fatalf("could not get file requests: %v", err)
		return nil, err
	}
	log.Printf("File Requests: %s", marketServerResponse.GetRequests())
	return marketServerResponse.GetRequests(), nil
}
