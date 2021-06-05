package util

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = bcrypt.DefaultCost
)

func GenerateBCryptoHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
