package repository

import "github.com/yasszu/go-jwt-auth/domain/entity"

//go:generate mockgen -source=./account.go -destination=./mock/account.go -package=mock
type AccountRepository interface {
	GetAccountByEmail(email string) (*entity.Account, error)
	GetAccountByID(accountID uint) (*entity.Account, error)
	CreateAccount(account *entity.Account) error
	UpdateAccount(account *entity.Account) error
	DeleteAccount(accountID uint) error
}
