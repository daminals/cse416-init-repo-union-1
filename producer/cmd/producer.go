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

	// Set up a connection to the market server check for file requests for each file hash
	for fileHash := range internal.FileHashes {
		fileRequests, err := internal.GetFileRequests(*marketServerAddr, fileHash)
		if err != nil {
			panic(err)
		}

		// Send the file link to each consumer requesting the file
		for _, fileRequest := range fileRequests {
			internal.SendFileLink(fileRequest.GetIp(), uint16(fileRequest.GetPort()), fileHash)
		}
	}

	// Adds a sample access token for testing
	internal.AccessTokens["123"] = &internal.ConsumerRequestInfo{
		RequestedFileHash: "file123Hash456",
		NumSentChunks:     0,
	}
}
