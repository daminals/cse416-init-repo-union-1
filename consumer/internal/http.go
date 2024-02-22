package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"
)

const ConsumerAddr string = "127.0.0.1" // this is the source ip address, change it to 0.0.0.0 in production
const ConsumerPort uint16 = 50052       // this is the source port, can be anything (should be recorded in market)

const token string = "" // this is the access token for the consumer

func GetFile(producer *pb.FileProducer) ([]byte, error) {
	// check if the file link is empty
	if producer.GetLink() == "" || producer.GetPaymentAddress() == "" {
		return nil, fmt.Errorf("no file link present")
	}

	// Check if the link contains the protocol
	fileURL := producer.GetLink()
	if !strings.HasPrefix(fileURL, "http://") && !strings.HasPrefix(fileURL, "https://") {
		fileURL = "http://" + fileURL
	}

	// send an http request to the producer to download the file
	netClient := &http.Client{}
	req, err := http.NewRequest("GET", fileURL, nil)
	if err != nil {
		log.Fatalf("Error creating http request: %v", err)
	}

	// add the access token in the header
	req.Header.Set("Authorization", "Bearer "+token)

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
