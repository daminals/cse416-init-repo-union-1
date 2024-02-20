package main

import (
	"github.com/daminals/cse416-init-repo-union-1/producer/internal"
)

// var (
// 	addr = flag.String("addr", "localhost:50051", "the address to connect to")
// )

func main() {
	// flag.Parse()
	// // Set up a connection to the market server check for file requests
	// fileRequests, err := internal.MarketServerRequest(*addr)
	// if err != nil {
	// 	for _, fileRequest := range fileRequests {
	// 		consumerAddr := fmt.Sprintf("%s:%d", fileRequest.GetIp(), fileRequest.GetPort())
	// 		internal.FileRequests[consumerAddr] = internal.Consumer{Addr: consumerAddr}
	// 	}
	// }

	// Adds localhost as a consumer for testing
	internal.FileRequests["127.0.0.1"] = &internal.Consumer{IPAddress: "127.0.0.1"}

	internal.StartServer()
}
