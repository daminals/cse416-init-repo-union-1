package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ProducerWallet = "wallet_address"

var (
	marketServerConnection *grpc.ClientConn
)

func StartMarketServerConnection(marketServerAddr string) (error) {
	// Set up a connection to the market server and add name to the list of producers per hash
	log.Printf("Connecting to market server at %s...", marketServerAddr)
	connMarketServer, err := grpc.Dial(marketServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	marketServerConnection = connMarketServer
	return nil
}

func CloseMarketServerConnection(marketServerAddr string) {
	// TODO: can be connected to many market servers (distributed)
	// so should have map which represents every market server, and close the connection to the specified one
	marketServerConnection.Close()
	log.Printf("Closed connection to market server at %s", marketServerAddr)
}

func EnrollProducer(marketServerAddr, fileHash string) (error) {
	clientMarketServer := pb.NewMarketServiceClient(marketServerConnection)
	
	// create producer
	producer := &pb.FileProducer{
		Hash:           fileHash,
		Link:           fmt.Sprintf("%s/%s", ProducerAddr, fileHash),
		Price:          0,
		PaymentAddress: ProducerWallet,
	}

	// Contact the server and print out its response.
	ctxMarketServer, cancelMarketServer := context.WithTimeout(context.Background(), time.Second)
	defer cancelMarketServer()
	_, err := clientMarketServer.AddProducer(ctxMarketServer, producer)
	if err != nil {
		log.Fatalf("could not get file requests: %v", err)
		return err
	}
	log.Printf("Sent: file hash availability for %s to market server at %s", fileHash, marketServerAddr)
	return nil
}
