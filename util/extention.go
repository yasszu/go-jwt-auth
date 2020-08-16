package util

import (
	"crypto/sha256"
	"encoding/hex"
)

type Password string

func (s Password) SHA256() string {
	bytes := sha256.Sum256([]byte(s))
	hash := hex.EncodeToString(bytes[:])
	return hash
}
