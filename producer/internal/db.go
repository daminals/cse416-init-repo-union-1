package internal

import (
	"crypto/rand"
	"math/big"
)

type ConsumerRequestInfo struct {
	RequestedFileURL string
	NumSentChunks    uint16
}

var AccessTokens = make(map[string]*ConsumerRequestInfo)

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
