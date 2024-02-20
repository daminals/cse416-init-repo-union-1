package main

import (
	"context"
	"flag"
	"log"
	"time"

	"crypto/rand"
	"encoding/base64"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	// consumerURL = flag.String("consumer-url", "localhost:50052", "placeholder URL")
)

// Access token length
const accessTokenLength = 32

func generateAccessToken(length int) (string, error) {
	// Calculate the number of bytes needed to represent the random string
	numBytes := length / 4 * 3
	if length%4 != 0 {
		numBytes = (length/4 + 1) * 3
	}

	// Generate random bytes
	bytes := make([]byte, numBytes)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	token := base64.URLEncoding.EncodeToString(bytes)

	// Trim the string to the desired length
	token = token[:length]

	return token, nil
}

func main() {
	flag.Parse()
	// Set up a connection to the market server.
	connMarket, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to market server: %v", err)
	}
	defer connMarket.Close()
	marketClient := pb.NewMarketServiceClient(connMarket)

	// Contact the market server and print out its response.
	ctxMarket, cancelMarket := context.WithTimeout(context.Background(), time.Second)
	defer cancelMarket()
	fileRequests, err := marketClient.GetFileRequests(ctxMarket, &pb.FileHash{Hash: "hash"})
	if err != nil {
		log.Fatalf("could not get file requests from market server: %v", err)
	}
	log.Printf("File Requests from market server: %s", fileRequests.GetRequests())
	
	// loop through the file requests and send the file links to the consumer

	// Create a map to store file addresses with their corresponding access tokens
	fileTokenMap := make(map[string]string)

	// Send file addresses to the consumer
for _, fileAddress := range fileRequests.GetRequests() {
	consumerURL := fileAddress.Ip + ":" + string(fileAddress.Port)

	// Set up a connection to the consumer.
	connConsumer, err := grpc.Dial(consumerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to consumer: %v", err)
	}
	defer connConsumer.Close()
	consumerClient := pb.NewConsumerServiceClient(connConsumer)

	log.Printf("Connected to consumer: %s", consumerURL)

    // Generate access token for each file
    token, err := generateAccessToken(accessTokenLength)
    if err != nil {
        log.Printf("Error generating access token: %v", err)
        continue
    }

    // Store file address and access token in the map
    fileTokenMap[fileAddress.Ip] = token

    // Construct the file link object
    fileLink := &pb.FileLink{
        Link:            fileAddress.Ip,
        Token:           token,
        PaymentAddress:  "payment_address", // Placeholder for payment address
    }

    // Send the file link to the consumer
    ctxConsumer, cancelConsumer := context.WithTimeout(context.Background(), time.Second)
    defer cancelConsumer()

    // Send the file link object to the consumer
    response, err := consumerClient.ReceiveFileInfo(ctxConsumer, fileLink)
    if err != nil {
        log.Printf("Failed to send file address to consumer: %v", err)
        continue
    }
    log.Printf("Response from consumer for %s: %v", fileAddress.Ip, response)
}

	// Print the fileTokenMap
	log.Printf("File addresses with corresponding access tokens: %v", fileTokenMap)
}
