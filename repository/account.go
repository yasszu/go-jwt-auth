package repository

import (
	"go-jwt-auth/model"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountByEmail(email string) (*model.Account, error)
	GetAccountById(accountID uint) (*model.Account, error)
	CreateAccount(account *model.Account) error
	UpdateAccount(account *model.Account) error
	DeleteAccount(accountID uint) error
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) GetAccountByEmail(email string) (*model.Account, error) {
	var account model.Account
	err := r.db.Where("email = ?", email).First(&account).Error
	return &account, err
}

func (r *AccountRepository) GetAccountById(id uint) (*model.Account, error) {
	var account model.Account
	err := r.db.First(&account, id).Error
	return &account, err
}

func (r *AccountRepository) CreateAccount(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) UpdateAccount(account *model.Account) error {
	return r.db.Save(account).Error
}

func (r *AccountRepository) DeleteAccount(accountID uint) error {
	return r.db.Delete(&model.Account{}, accountID).Error
}
