package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
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
