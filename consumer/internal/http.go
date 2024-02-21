package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const ConsumerAddr string = "127.0.0.1" // this is the source ip address, change it to 0.0.0.0 in production
const ConsumerPort uint16 = 50052

func GetFile() ([]byte, error) {
	// check if the file link is empty
	if CurrentFileLink.Link == "" {
		return nil, fmt.Errorf("no file link present")
	}

	// send an http request to the producer to download the file
	netClient := &http.Client{}
	req, err := http.NewRequest("GET", CurrentFileLink.Link, nil)
	if err != nil {
		log.Fatalf("Error creating http request: %v", err)
	}

	// add the access token in the header
	req.Header.Set("Authorization", "Bearer "+CurrentFileLink.Token)

	// send the request
	res, err := netClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending http request: %v", err)
	}

	// check if the response is 200 OK
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Error downloading file: %v", res.Status)
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	return responseBody, nil
}
