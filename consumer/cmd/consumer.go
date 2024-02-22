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

	// send file request will ask the market server who has my file, and how much is it
	err := internal.SendFileRequest(*marketServerAddr)
	if err != nil {
		log.Fatalf("Failed to send file request: %v", err)
	}

	// Start gRPC server
	// internal.FileReceived.Add(1)
	// go internal.StartListener()

	// Send an http request to the producer to download the file
	// internal.FileReceived.Wait()
	// internal.StopListener()

	// add some mechanism for consumer to choose which producer to download from
	// for now, just download from the test producer


	// Get the file chunk from the producer
	file, err := internal.GetFile()
	if err != nil {
		log.Fatalf("Failed to get file: %v", err)
	}

	// Print the file
	log.Printf("Recieved: File %s", string(file))
}
