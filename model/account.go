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

type SignupForm struct {
	Username string `form:"username" validate:"required,max=40"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=64"`
}

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type AccountResponse struct {
	AccountID uint      `json:"account_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"crated_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (a *Account) Populate(form *SignupForm) error {
	hash, err := util.GenerateBCryptoHash(form.Password)
	if err != nil {
		return err
	}
	a.Username = form.Username
	a.Email = form.Email
	a.PasswordHash = hash
	return err
}

func NewAccountResponse(a *Account) AccountResponse {
	return AccountResponse{
		AccountID: a.ID,
		Username:  a.Username,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
	}
}
