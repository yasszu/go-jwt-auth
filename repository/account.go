package repository

import (
	"go-jwt-auth/model"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountByEmail(email string) (*model.Account, error)
	GetAccountById(id uint) (*model.Account, error)
	CreateAccount(account *model.Account) error
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (a *AccountRepository) GetAccountByEmail(email string) (*model.Account, error) {
	var account model.Account
	err := a.db.Where("email = ?", email).First(&account).Error
	return &account, err
}

func (a *AccountRepository) GetAccountById(id uint) (*model.Account, error) {
	var account model.Account
	err := a.db.Where("account_id = ?", id).First(&account).Error
	return &account, err
}

func (a *AccountRepository) CreateAccount(account *model.Account) error {
	err := a.db.Create(account).Error
	return err
}
