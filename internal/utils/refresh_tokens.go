package utils

import (
	"cinema-service/internal/configs/structures"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func ParseRefreshToken(tokenStr string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return structures.RefreshJwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
