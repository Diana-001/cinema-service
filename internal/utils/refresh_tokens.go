package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var refreshSecret = []byte("super_secret_refresh")

type RefreshClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func ParseRefreshToken(tokenStr string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
