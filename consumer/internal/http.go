package internal

import (
	"net/http"
)

func HandleFileRequest(writer http.ResponseWriter, request *http.Request) {
	// fileHash := mux.Vars(request)["fileHash"]
	// accessToken := request.Header.Get("Authorization")
	// consumerAddr := request.RemoteAddr

	writer.Write([]byte("File Requested!"))
}
