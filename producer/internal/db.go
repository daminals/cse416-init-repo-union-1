package internal

import (
	"crypto/rand"
	"math/big"
)

type ConsumerRequestInfo struct {
	RequestedFileHash string
	NumSentChunks     uint16
}

// Map of active access tokens to the consumer that recieved them
var AccessTokens = make(map[string]*ConsumerRequestInfo)

// Set of file hashes that the producer offers
var FileHashes = make(map[string]bool)

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
