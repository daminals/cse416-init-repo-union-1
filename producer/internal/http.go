package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const ProducerAddr = "127.0.0.1:8080/file"

func HandleFileRequest(writer http.ResponseWriter, request *http.Request) {
	fileHash := mux.Vars(request)["fileHash"]
	accessToken := request.Header.Get("Authorization")
	if accessToken != "" {
		_, accessToken, _ = strings.Cut(accessToken, "Bearer ")
	}
	consumerAddr, _, _ := strings.Cut(request.RemoteAddr, ":")

	// Verifies the consumer's IP address is expected according to the market server
	consumer, ok := FileRequests[consumerAddr]
	if !ok {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte(consumerAddr))
		log.Printf("Consumer not found: %s", consumerAddr)
		return
	}

	// Verifies the request contains the same file hash as previous requests
	if consumer.RequestedFileHash == "" {
		consumer.RequestedFileHash = fileHash
	} else if consumer.RequestedFileHash != fileHash {
		writer.WriteHeader(http.StatusUnauthorized)
		log.Printf("Consumer requested file: %s, expected file: %s", fileHash, consumer.RequestedFileHash)
		return
	}

	// Verifies the consumer's access token is valid
	if (consumer.NumReceievedChunks == 0 && consumer.NextAccessToken == "") || consumer.NextAccessToken == accessToken {
		consumer.NextAccessToken = GenerateAccessToken()
		log.Printf("Consumer given access token (%d): %s", consumer.NumReceievedChunks, consumer.NextAccessToken)
	} else if consumer.NextAccessToken != accessToken {
		writer.WriteHeader(http.StatusUnauthorized)
		log.Printf("Consumer provided invalid access token: %s, expected %s", accessToken, consumer.NextAccessToken)
		return
	}

	// Placeholder for the file chunk
	writer.Write([]byte("File Chunk\n"))
	consumer.NumReceievedChunks++

	// Sends the consumer info back for debugging
	consumerJson, _ := json.Marshal(*consumer)
	writer.Write(consumerJson)
}

func StartServer() {
	server := &http.Server{Addr: ":8080"}

	router := mux.NewRouter()
	router.HandleFunc("/file/{fileHash}", HandleFileRequest).Methods("GET")
	server.Handler = router
	server.ListenAndServe()
}
