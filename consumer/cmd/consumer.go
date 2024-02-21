package main

import (
	"flag"
	"log"

	"github.com/daminals/cse416-init-repo-union-1/consumer/internal"
)

var (
	marketServerAddr = flag.String("market-server-address", "127.0.0.1:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	err := internal.SendFileRequest(*marketServerAddr)
	if err != nil {
		log.Fatalf("Failed to send file request: %v", err)
	}

	// Start gRPC server
	internal.StartListener()

	// Send an http request to the producer to download the file
	file, err := internal.GetFile()
	if err != nil {
		log.Fatalf("Failed to get file: %v", err)
	}

	// Print the file
	log.Printf("File: %s", string(file))
}
