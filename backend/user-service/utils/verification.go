package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"regexp"
	"strconv"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"golang.org/x/crypto/argon2"
)

func Encode(code string, current_time int64) string {
	config, _ := config.LoadConfig("..")
	passwordWithPepper := append([]byte(code), []byte(config.EmailCode)...)
	hashedPassword := argon2.IDKey(passwordWithPepper, []byte(strconv.FormatInt(current_time, 10)), argon2Time, argon2Memory, argon2Threads, argon2KeyLen)
	b64Hash := customBase64EncodeToString(hashedPassword)
	return b64Hash
}

func Compare(sid string, current_time int64, candidateCode string) bool {
	b64Hash := Encode(sid, current_time)
	if subtle.ConstantTimeCompare([]byte(b64Hash), []byte(candidateCode)) != 1 {
		return false
	}
	return true
}

func customBase64EncodeToString(input []byte) string {
	const customBase64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	encoded := make([]byte, base64.RawStdEncoding.EncodedLen(len(input)))

	base64.RawStdEncoding.Encode(encoded, input)

	regExp := regexp.MustCompile("[^a-zA-Z0-9]")
	result := regExp.ReplaceAllString(string(encoded), "")

	return result
}
