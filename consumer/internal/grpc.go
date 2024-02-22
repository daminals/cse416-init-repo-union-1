package internal

import (
	"context"
	"log"
	"time"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	hash 					string = "hash"
)

func GetProducerList(marketServerAddr string) ([]*pb.FileProducer, error) {
	// Set up a connection to the server.
	log.Printf("Connecting to market server at %s...", marketServerAddr)
	connMarketServer, err := grpc.Dial(marketServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error: did not connect to market server: %v", err)
		return nil, err
	}
	clientMarketServer := pb.NewMarketServiceClient(connMarketServer)
	defer connMarketServer.Close() // close connection after function ends

	// Contact the server and print out its response.
	ctxMarketServer, cancelMarketServer := context.WithTimeout(context.Background(), time.Second)
	defer cancelMarketServer() // cancel the context after function ends

	// get all the producers who are selling the file hash
	producers, err := clientMarketServer.GetProducers(ctxMarketServer, &pb.FileHash{Hash: hash})
	if err != nil {
		log.Printf("Error: could not get producers: %v", err)
		return nil, err
	}
	return producers.GetProducers(), nil
}
