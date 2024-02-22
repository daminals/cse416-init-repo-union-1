package internal

import (
	"crypto/rand"
	"math/big"
)

type ConsumerRequestInfo struct {
	NumSentChunks uint64
	AccessToken   string
}

type Consumer struct {
	Requests      map[string]*ConsumerRequestInfo
	WalletAddress string
}

// Map of active access tokens to the consumer that recieved them
var Consumers = make(map[string]*Consumer)

// Set of file hashes that the producer offers with their sizes
var FileHashes = make(map[string]uint64)

func GenerateAccessToken() string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+=@#")

	charsLen := big.NewInt(int64(len(chars)))

	accessToken := make([]rune, 64)
	for i := range accessToken {
		randomIndex, _ := rand.Int(rand.Reader, charsLen)
		accessToken[i] = chars[randomIndex.Int64()]
	}

	return string(accessToken)
}
