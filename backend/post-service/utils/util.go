package utils

import (
	"math/rand"
	"time"
)

func IsElementInStringArray(s string, arr *[]string) bool {
	for _, a := range *arr {
		if a == s {
			return true
		}
	}

	return false
}

func IsItemInArray(s int64, arr *[]int64) bool {
	for _, a := range *arr {
		if a == s {
			return true
		}
	}

	return false
}

func CheckArrayEqual(s1 *[]string, s2 *[]string) bool {
	if len(*s1) != len(*s2) {
		return false
	}

	for _, s := range *s1 {
		if IsElementInStringArray(s, s2) == false {
			return false
		}
	}

	return true
}

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
