package internal

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const ProducerAddr = "127.0.0.1:8080"

func HandleFileRequest(writer http.ResponseWriter, request *http.Request) {
	// Path Variables
	fileHash := mux.Vars(request)["fileHash"]
	requestedChunkIndex := mux.Vars(request)["fileChunkIndex"]

	// Header Data
	consumerAddr, _, _ := strings.Cut(request.RemoteAddr, ":")
	accessToken := request.Header.Get("Authorization")
	if accessToken != "" {
		_, accessToken, _ = strings.Cut(accessToken, "Bearer ")
	}

	consumer, ok := Consumers[consumerAddr]
	if !ok {
		// If consumer IP not found, create a new consumer
		Consumers[consumerAddr] = &Consumer{
			WalletAddress: "wallet_address", // Placeholder for the consumer's wallet address
			Requests:      make(map[string]*ConsumerRequestInfo),
		}
		consumer = Consumers[consumerAddr]
	}

	requestInfo, ok := consumer.Requests[fileHash]
	if !ok {
		// If file hash not found, create a new request info
		consumer.Requests[fileHash] = &ConsumerRequestInfo{
			AccessToken:   "",
			NumSentChunks: 0,
		}
		requestInfo = consumer.Requests[fileHash]
	}

	// Check if the access token is expected (first request is free, accessToken is "")
	if requestInfo.AccessToken != accessToken {
		writer.WriteHeader(http.StatusUnauthorized)
		log.Printf("Unauthorized: consumer at %s sent an invalid access token. Expected %s and got %s", request.RemoteAddr, requestInfo.AccessToken, accessToken)
		return
	}

	// Check if the consumer is requesting a specific chunk
	chunkIndex, err := strconv.ParseUint(requestedChunkIndex, 10, 64)
	if err != nil {
		chunkIndex = requestInfo.NumSentChunks
	}

	// Placeholder for the file chunk
	writer.Write([]byte("File Chunk of " + fileHash))
	requestInfo.NumSentChunks++
	log.Printf("Sent: file chunk %d of %s to consumer at %s", chunkIndex, fileHash, request.RemoteAddr)

	// Generate a new access token to be used for the next request
	newAccessToken := GenerateAccessToken()
	requestInfo.AccessToken = newAccessToken
	log.Printf("Generated new access token (%s) for file (%s) for consumer (%s)", newAccessToken, fileHash, request.RemoteAddr)
}

func StartServer() {
	server := &http.Server{Addr: ":8080"}

	router := mux.NewRouter()
	router.HandleFunc("/{fileHash}", HandleFileRequest).Methods("GET")
	router.HandleFunc("/{fileHash}/{fileChunkIndex}", HandleFileRequest).Methods("GET")
	server.Handler = router
	server.ListenAndServe()
}
