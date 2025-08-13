package utils

import (
	"math/rand"
	"time"
)

var length int = 30 // default length for the API key

func Generate() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$^&*()-_=+[]{}.?/"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	APIKEY := make([]byte, length)
	for i := range APIKEY {
		APIKEY[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(APIKEY)
}
