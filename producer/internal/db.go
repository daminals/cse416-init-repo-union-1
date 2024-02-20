package internal

import (
	"crypto/rand"
	"math/big"
)

type Consumer struct {
	NumReceievedChunks uint16
	IPAddress          string
	RequestedFileHash  string
	NextAccessToken    string
}

var FileRequests = make(map[string]*Consumer)

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
