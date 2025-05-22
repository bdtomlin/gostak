package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("must be at least 8 characters")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
