package entity

import "time"

type AccessToken struct {
	AccountID uint
	Token     string
	ExpiresAt time.Time
}

type TokenPair struct {
	AccessToken  *AccessToken
	RefreshToken *AccessToken
}
