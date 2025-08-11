package utils

import (
	"split-bill-service/config"
	"split-bill-service/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	User *models.User `json:"user"`
	jwt.RegisteredClaims
}

var secretKey = config.GetEnv("JWT_SECRET", "secret")

func GenerateJWTUserToken(user *models.User, duration time.Duration) (string, error) {
	claims := &TokenClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWTUserToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return token.Claims.(*TokenClaims), nil
}

