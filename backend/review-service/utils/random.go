package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	var result string
	charsetLength := len(charset)

	for i := 0; i < length; i++ {
		randomIndex := randomGenerator.Intn(charsetLength)
		result += string(charset[randomIndex])
	}

	return result
}

func GenerateRandomNumber(length int) string {
	charset := "0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	var result string
	charsetLength := len(charset)

	for i := 0; i < length; i++ {
		randomIndex := randomGenerator.Intn(charsetLength)
		result += string(charset[randomIndex])
	}

	return result
}