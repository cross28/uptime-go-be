package auth

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)

	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return hash, nil
}

func VerifyPassword(attempted_password string, actual_password_hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(attempted_password, actual_password_hash)
	if err != nil {
		return false
	}
	return match
}