package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "d6c0f5f6-9b50-4dc3-bb78-75a2780c5b0a"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // each token expires after 2 hours
	})

	return token.SignedString([]byte(secretKey))
}

func ReadToken(token jwt.Token) {
	//jwt.
}
