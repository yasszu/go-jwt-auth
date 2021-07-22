package entity

import (
	"time"
)

type Account struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AccountResponse struct {
	AccountID uint   `json:"account_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func (e *Account) Response() AccountResponse {
	return AccountResponse{
		AccountID: e.ID,
		Username:  e.Username,
		Email:     e.Email,
	}
}
