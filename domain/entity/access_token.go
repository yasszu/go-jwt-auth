package entity

import "time"

type AccessToken struct {
	AccountID uint
	Token     string
	ExpiresAt time.Time
}
