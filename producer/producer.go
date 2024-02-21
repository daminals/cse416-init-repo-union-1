package main

import (
	"context"
	"flag"
	"log"
	"time"

	"crypto/rand"
	"encoding/base64"
	"strconv"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

const accessTokenLength = 32 // Length of the access token

func generateAccessToken(existingTokens map[string]string) (string, error) {
	for {
		// number of bytes needed to represent the random string
		numBytes := accessTokenLength / 4 * 3
		if accessTokenLength%4 != 0 {
			numBytes = (accessTokenLength/4 + 1) * 3
		}

		// generate random bytes
		bytes := make([]byte, numBytes)
		if _, err := rand.Read(bytes); err != nil {
			return "", err
		}

		// Encode the random bytes to base64
		token := base64.URLEncoding.EncodeToString(bytes)

		// Trim the string to the desired length
		token = token[:accessTokenLength]

		// Check if the token already exists
		if _, exists := existingTokens[token]; !exists {
			// Token is unique, return it
			return token, nil
		}
	}
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
	log.Printf("Received: file requests from market server: %s", fileRequests.GetRequests())

	// Create a map to store file addresses with their corresponding access tokens
	fileTokenMap := make(map[string]string)

	// loop through the file requests and send the file links to the consumer
	for _, fileAddress := range fileRequests.GetRequests() {
		consumerAddr := fileAddress.Ip + ":" + strconv.Itoa(int(fileAddress.Port))
		log.Printf("Sent: file address to consumer at %s", consumerAddr)

		// Set up a connection to the consumer.
		connConsumer, err := grpc.Dial(consumerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect to consumer: %v", err)
		}
		consumerClient := pb.NewConsumerServiceClient(connConsumer)

		// Generate access token for each file
		token, err := generateAccessToken(fileTokenMap)
		if err != nil {
			log.Printf("Error generating access token: %v", err)
			continue
		}

		// Store file address and access token in the map
		fileTokenMap[fileAddress.Ip] = token

		// Construct the file link object
		fileLink := &pb.FileLink{
			Link:           fileAddress.Ip,
			Token:          token,
			PaymentAddress: "payment_address", // Placeholder for payment address
		}

		// Send the file link to the consumer
		ctxConsumer, cancelConsumer := context.WithTimeout(context.Background(), time.Second)

		// Send the file link object to the consumer
		response, err := consumerClient.ReceiveFileInfo(ctxConsumer, fileLink)
		cancelConsumer() // Cancel context after the RPC call
		if err != nil {
			log.Printf("Failed to send file address to consumer: %v", err)
			continue
		}
		log.Printf("Recieved: %v from %s", response, fileAddress.Ip)
		// Close the connection to the consumer
		connConsumer.Close()
	}

	// Print the fileTokenMap
	// log.Printf("File addresses with corresponding access tokens: %v", fileTokenMap)
}
