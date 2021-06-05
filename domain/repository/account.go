package repository

import "go-jwt-auth/domain/entity"

type AccountRepository interface {
	GetAccountByEmail(email string) (*entity.Account, error)
	GetAccountByID(accountID uint) (*entity.Account, error)
	CreateAccount(account *entity.Account) error
	UpdateAccount(account *entity.Account) error
	DeleteAccount(accountID uint) error
}
