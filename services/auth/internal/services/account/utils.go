package account

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func passwordEncode(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}
	return string(hashPass), nil
}

func passwordCompare(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("bcrypt.CompareHashAndPassword: %w", err)
	}

	return nil
}
