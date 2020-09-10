package model

import (
	"go-jwt-auth/util"
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

type AccountForm struct {
	Username string `form:"username" validate:"required,max=40"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=64"`
}

type AccountResponse struct {
	AccountID uint      `json:"account_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"crated_at"`
}

func (a *Account) Populate(form *AccountForm) error {
	hash, err := util.GenerateBCryptoHash(form.Password)
	if err != nil {
		return err
	}
	a.Username = form.Username
	a.Email = form.Email
	a.PasswordHash = hash
	return nil
}

func NewAccountResponse(a *Account) AccountResponse {
	return AccountResponse{
		AccountID: a.ID,
		Username:  a.Username,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
	}
}
