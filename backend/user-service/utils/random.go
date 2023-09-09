package utils

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func GenerateRandomNumber(length int) (string, error) {
	randomBytes := make([]byte, (length+1)/2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomInt := new(big.Int)
	randomInt.SetBytes(randomBytes)

	randomNumber := fmt.Sprintf("%0*d", length, randomInt)
	return randomNumber, nil
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
