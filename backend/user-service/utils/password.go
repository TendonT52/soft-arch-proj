package utils

import (
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
	config, _ := config.LoadConfig("..")
	// Decode the hash password (get from db) from base64
	storedHashBytes, err := base64.RawStdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return err
	}

	// Prepare the user's password with the pepper
	passwordWithPepper := append([]byte(candidatePassword), []byte(config.Pepper)...)

	// Hash the user's password with the same parameters
	hashedCandidatePassword := argon2.IDKey(
		passwordWithPepper,
		[]byte(strconv.FormatInt(current_time, 10)),
		argon2Time,
		argon2Memory,
		argon2Threads,
		argon2KeyLen,
	)

	// Compare the newly hashed password with the stored hash
	result := compareHashes(storedHashBytes, hashedCandidatePassword)
	if !result {
		return errors.New("Password is not correct")
	}
	return nil
}

// CompareHashes compares two byte slices in constant time to avoid timing attacks.
func compareHashes(hash1, hash2 []byte) bool {
	if len(hash1) != len(hash2) {
		return false
	}

	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}

	return true
}