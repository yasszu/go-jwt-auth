package repository

import (
	"context"

	"github.com/yasszu/go-jwt-auth/domain/entity"
)

//go:generate mockgen -source=./account.go -destination=./mock/account.go -package=mock
type Account interface {
	GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error)
	GetAccountByID(ctx context.Context, accountID uint) (*entity.Account, error)
	CreateAccount(ctx context.Context, account *entity.Account) error
	UpdateAccount(ctx context.Context, account *entity.Account) error
	DeleteAccount(ctx context.Context, accountID uint) error
}
