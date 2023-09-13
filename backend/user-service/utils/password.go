package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strconv"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"golang.org/x/crypto/argon2"
)

const (
	argon2Time    = 1
	argon2Memory  = 64 * 1024
	argon2Threads = 4
	argon2KeyLen  = 32
)

func HashPassword(password string, current_time int64) string {
	config, _ := config.LoadConfig("..")
	passwordWithPepper := append([]byte(password), []byte(config.Pepper)...)
	hashedPassword := argon2.IDKey(passwordWithPepper, []byte(strconv.FormatInt(current_time, 10)), argon2Time, argon2Memory, argon2Threads, argon2KeyLen)
	b64Hash := base64.RawStdEncoding.EncodeToString(hashedPassword)
	return b64Hash
}

func VerifyPassword(hashedPassword string, candidatePassword string, current_time int64) error {
	b64Hash := HashPassword(candidatePassword, current_time)
	if subtle.ConstantTimeCompare([]byte(hashedPassword), []byte(b64Hash)) != 1 {
		return errors.New("password not match")
	}
	return nil
}
