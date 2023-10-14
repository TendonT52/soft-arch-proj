package mock

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/JinnnDamanee/review-service/config"
	"github.com/JinnnDamanee/review-service/domain"
	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(ttl time.Duration, payload *domain.Payload) (string, error) {
	config, err := config.LoadConfigTest("..")
	if err != nil {
		return "", fmt.Errorf("create: load config: %w", err)
	}
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(config.AccessTokenPrivateKeyTest)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["userId"] = payload.UserId
	claims["role"] = payload.Role
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	// jwt.SigningMethodHS256()
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("create: sign access-token: %w", err)
	}

	return token, nil
}
