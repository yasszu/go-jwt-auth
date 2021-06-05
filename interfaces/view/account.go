package view

import (
	"go-jwt-auth/domain/entity"
)

type Account struct {
	AccountID uint   `json:"account_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func NewAccount(a *entity.Account) Account {
	return Account{
		AccountID: a.ID,
		Username:  a.Username,
		Email:     a.Email,
	}
}
