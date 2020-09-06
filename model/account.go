package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
}

type AccountForm struct {
	Username string `validate:"required,max=40"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=64"`
}

type AccountResponse struct {
	AccountID uint      `json:"account_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"crated_at"`
}

func NewAccountResponse(a *Account) AccountResponse {
	return AccountResponse{
		AccountID: a.ID,
		Username:  a.Username,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
	}
}
