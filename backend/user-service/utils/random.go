package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(length int) string {
	charset := "0123456789"
	rand.Seed(time.Now().UnixNano())

	var result string
	charsetLength := len(charset)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charsetLength)
		result += string(charset[randomIndex])
	}

	return result
}

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	var result string
	charsetLength := len(charset)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charsetLength)
		result += string(charset[randomIndex])
	}

	return result
}
