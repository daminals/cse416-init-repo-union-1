package main

import (
	"flag"

	"github.com/daminals/cse416-init-repo-union-1/producer/internal"
)

var (
	marketServerAddr = flag.String("market-server-address", "127.0.0.1:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// Starts the HTTP server on another process
	go internal.StartServer()

	hash := "hash" // could be any hash i have

	// Set up a connection to the market server check for file requests
	fileRequests, err := internal.GetFileRequests(*marketServerAddr, hash)
	if err != nil {
		panic(err)
	}

	// Send the file link to each consumer
	for _, fileRequest := range fileRequests {
		internal.FileRequests[fileRequest.GetIp()] = &internal.Consumer{
			IPAddress:          fileRequest.GetIp(),
			NextAccessToken:    internal.GenerateAccessToken(),
			RequestedFileHash:  hash, // Should be provided by the market server
			NumReceievedChunks: 0,
		}

		internal.SendFileLink(fileRequest.GetIp(), uint16(fileRequest.GetPort()), hash) // Should be provided by the market server
	}

	// Adds localhost as a consumer for testing
	internal.FileRequests["127.0.0.1"] = &internal.Consumer{
		IPAddress:          "127.0.0.1",
		NextAccessToken:    internal.GenerateAccessToken(),
		RequestedFileHash:  "123", // Should be provided by the market server
		NumReceievedChunks: 0,
	}
}
