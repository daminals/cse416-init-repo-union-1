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
	producers, err := internal.GetProducerList(*marketServerAddr)
	if err != nil {
		log.Fatalf("Failed to send file request: %v", err)
	}

	// define some kind of mechanism to select the producer
	// e.g. select the producer with the lowest price
	// for now, just select the first producer
	producer := producers[0]

	// Get the file chunk from the producer
	file, err := internal.GetFile(producer)
	if err != nil {
		log.Fatalf("Failed to get file: %v", err)
	}

	// Print the file
	log.Printf("Recieved: File %s", string(file))
}
