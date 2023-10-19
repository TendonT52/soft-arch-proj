package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/JinnnDamanee/review-service/config"
	"github.com/JinnnDamanee/review-service/domain"
	"github.com/golang-jwt/jwt"
)

func ValidateAccessToken(token string) (*domain.Payload, error) {
	config, _ := config.LoadConfig("..")
	decodedPublicKey, err := base64.StdEncoding.DecodeString(config.AccessTokenPublicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	payload := &domain.Payload{
		UserId: int64(claims["userId"].(float64)),
		Role:   claims["role"].(string),
	}

	return payload, nil
}
