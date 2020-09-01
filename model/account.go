package model

import "time"

type Account struct {
	AccountID uint `gorm:"primary_key"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
		AccountID: a.AccountID,
		Username:  a.Username,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
	}
}
