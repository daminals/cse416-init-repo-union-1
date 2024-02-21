package internal

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const ProducerAddr = "127.0.0.1:8080"

func HandleFileRequest(writer http.ResponseWriter, request *http.Request) {
	fileHash := mux.Vars(request)["fileHash"]
	accessToken := request.Header.Get("Authorization")
	if accessToken == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		log.Printf("No access token provided by consumer (%s)", request.RemoteAddr)
		return
	} else {
		_, accessToken, _ = strings.Cut(accessToken, "Bearer ")
	}

	requestInfo, ok := AccessTokens[accessToken]
	if !ok {
		writer.WriteHeader(http.StatusUnauthorized)
		log.Printf("Error: Consumer (%s) provided invalid access token %s", request.RemoteAddr, accessToken)
		return
	}

	// Placeholder for the file chunk
	writer.Write([]byte("File Chunk of " + fileHash))
	requestInfo.NumSentChunks++
	log.Printf("Sent: file chunk %d of %s to consumer at %s", requestInfo.NumSentChunks, fileHash, request.RemoteAddr)

	// Generate a new access token to be used for the next request
	delete(AccessTokens, accessToken)
	newAccessToken := GenerateAccessToken()
	AccessTokens[newAccessToken] = requestInfo
	log.Printf("Sent: new access token (%s) to consumer at %s", newAccessToken, request.RemoteAddr)
}

func StartServer() {
	server := &http.Server{Addr: ":8080"}

	router := mux.NewRouter()
	router.HandleFunc("/{fileHash}", HandleFileRequest).Methods("GET")
	server.Handler = router
	server.ListenAndServe()
}
