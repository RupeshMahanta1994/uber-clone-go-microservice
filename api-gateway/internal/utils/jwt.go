package utils

import "github.com/golang-jwt/jwt/v5"

var JWTSecret = []byte("supersecret")

type Claims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenStr string) (*Claims, error) {
	token, error := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if error != nil {
		return nil, error
	}
	return token.Claims.(*Claims), nil
}
