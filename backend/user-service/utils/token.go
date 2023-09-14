package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(ttl time.Duration, payload *domain.Payload) (string, error) {
	config, _ := config.LoadConfig("..")
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(config.AccessTokenPrivateKey)
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

func CreateRefreshToken(ttl time.Duration, userId int64) (string, error) {
	config, _ := config.LoadConfig("..")
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.RefreshTokenSecret))
	if err != nil {
		return "", fmt.Errorf("create: sign refresh-token: %w", err)
	}

	return tokenString, nil
}

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

func ValidateRefreshToken(refreshToken string) (int64, error) {
	config, _ := config.LoadConfig("..")

	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.RefreshTokenSecret), nil
	})

	if err != nil {
		return 0, fmt.Errorf("verify: parse refresh-token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return 0, fmt.Errorf("verify: could not extract claims")
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return 0, fmt.Errorf("verify: userId claim is missing or invalid")
	}

	return int64(userID), nil
}
