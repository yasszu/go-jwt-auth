package jwt

import (
	"errors"
)

var (
	ErrorParseClaims  = errors.New("couldn't parse claims")
	ErrorTokenExpired = errors.New("jWT is expired")
)
