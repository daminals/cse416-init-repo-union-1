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
	marketServers map[string]*grpc.ClientConn = make(map[string]*grpc.ClientConn)
)

func StartMarketServerConnection(marketServerAddr string) (error) {
	// Set up a connection to the market server and add name to the list of producers per hash
	log.Printf("Connecting to market server at %s...", marketServerAddr)
	marketServerConn, err := grpc.Dial(marketServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	// add connection to map
	marketServers[marketServerAddr] = marketServerConn
	return nil
}

func CloseMarketServerConnection(marketServerAddr string) {
	// find the connection in the map and close it
	marketServerConnection, ok := marketServers[marketServerAddr]
	if !ok {
		log.Printf("Error: Connection to market server at %s not found, cannot close", marketServerAddr)
		return
	}
	marketServerConnection.Close()
	log.Printf("Closed connection to market server at %s", marketServerAddr)
}

func EnrollProducer(marketServerAddr, fileHash string) (error) {
	marketServerConnection, ok := marketServers[marketServerAddr]
	if !ok {
		log.Printf("Error: Connection to market server at %s not found, cannot enroll producer", marketServerAddr)
		return fmt.Errorf("Error: Connection to market server at %s not found, enroll producer", marketServerAddr)
	}
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
