package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	bytes, err := bcrypt.GenerateFromPassword(passwordBytes, 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func DoesHashMatchTerm(term string, hashedValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(term))
	if err == nil {
		return true
	} else {
		return false
	}
}
