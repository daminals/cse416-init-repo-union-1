package main

import (
	"flag"
	"sync"

	"github.com/daminals/cse416-init-repo-union-1/producer/internal"
)

var (
	marketServerAddr = flag.String("market-server-address", "127.0.0.1:50051", "the address to connect to")
	wg               sync.WaitGroup
)

func main() {
	flag.Parse()

	// Add the sample hash
	internal.FileHashes["hash"] = true

	// Starts the HTTP server on another process
	wg.Add(1)
	go func() {
		internal.StartServer()
		wg.Done()
	}()

	// Set up a connection to the market server
	err := internal.StartMarketServerConnection(*marketServerAddr)
	if err != nil {
		panic(err)
	}

	// Add my name to the list of producers per hash
	for fileHash := range internal.FileHashes {
		err := internal.EnrollProducer(*marketServerAddr, fileHash)
		if err != nil {
			panic(err)
		}
	}

	// Close the connection to the market server
	internal.CloseMarketServerConnection(*marketServerAddr)

	// Adds a sample access token for testing
	internal.AccessTokens["123"] = &internal.ConsumerRequestInfo{
		RequestedFileHash: "file123Hash456",
		NumSentChunks:     0,
	}

	wg.Wait()
}
