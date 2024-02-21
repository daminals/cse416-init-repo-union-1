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

func GetFileRequests(marketServerAddr, hash string) ([]*pb.FileRequest, error) {
	// Establish connection with the market server
	log.Printf("Connecting to market server at %s...", marketServerAddr)
	connMarketServer, err := grpc.Dial(marketServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer connMarketServer.Close()
	clientMarketServer := pb.NewMarketServiceClient(connMarketServer)

	// Contact the server and print out its response.
	ctxMarketServer, cancelMarketServer := context.WithTimeout(context.Background(), time.Second)
	defer cancelMarketServer()
	resMarketServer, err := clientMarketServer.GetFileRequests(ctxMarketServer, &pb.FileHash{Hash: hash})
	if err != nil {
		log.Fatalf("could not get file requests: %v", err)
		return nil, err
	}
	log.Printf("Received: file requests %s from market at %s", resMarketServer.GetRequests(), marketServerAddr)
	return resMarketServer.GetRequests(), nil
}

func SendFileLink(consumerAddr string, consumerPort uint16, fileHash string) {
	// Get the consumer info
	consumer, ok := FileRequests[consumerAddr]
	if !ok {
		log.Printf("Consumer not found: %s", consumerAddr)
		return
	}
	// Set up a connection to the consumer.
	fullConsumerAddr := fmt.Sprintf("%s:%d", consumerAddr, consumerPort)
	log.Printf("Connecting to consumer server at %s...", fullConsumerAddr)
	connConsumer, err := grpc.Dial(fullConsumerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to consumer: %v", err)
	}
	defer connConsumer.Close()
	clientConsumer := pb.NewConsumerServiceClient(connConsumer)

	// Create the file link to be sent
	fileLink := &pb.FileLink{
		Link:           fmt.Sprintf("%s/%s", ProducerAddr, fileHash),
		Token:          consumer.NextAccessToken,
		PaymentAddress: ProducerWallet,
	}

	// Send the file link to the consumer
	ctxConsumer, cancelConsumer := context.WithTimeout(context.Background(), time.Second)
	defer cancelConsumer()
	resConsumer, err := clientConsumer.ReceiveFileInfo(ctxConsumer, fileLink)
	if err != nil {
		log.Printf("Failed to send file address to consumer: %v", err)
		return
	}
	log.Printf("Recieved: %v from consumer at %s", resConsumer, fullConsumerAddr)
}
