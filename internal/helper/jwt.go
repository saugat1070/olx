package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func CreateAccessToken(payload JwtPayload, secretKey string) (string, error) {
	tokenTTL := 15 * time.Hour
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":       payload.ID,
		"id":        payload.ID,
		"email":     payload.Email,
		"full_name": payload.FullName,
		"role":      payload.Role,
		"exp":       now.Add(tokenTTL).Unix(),
		"iat":       now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	return signedToken, err
}
