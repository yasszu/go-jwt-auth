package entity

import "time"

type AccessToken struct {
	AccountID uint      `json:"account_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expired_at"`
}
