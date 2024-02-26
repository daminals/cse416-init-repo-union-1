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
	internal.FileHashes["hash"] = 1600000

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

	// Adds a sample Consumer entry for testing
	// internal.Consumers["127.0.0.1"] = &internal.Consumer{
	// 	Requests:      make(map[string]*internal.ConsumerRequestInfo),
	// 	WalletAddress: "wallet_address",
	// }
	// internal.Consumers["127.0.0.1"].Requests["hash"] = &internal.ConsumerRequestInfo{
	// 	AccessToken:   "",
	// 	NumSentChunks: 0,
	// }

	wg.Wait()
}
