package repository

import (
	"go-jwt-auth/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountByEmail(email string) (*model.Account, error)
	GetAccountById(accountID uint) (*model.Account, error)
	CreateAccount(account *model.Account) error
	UpdateAccount(account *model.Account) error
	DeleteAccount(accountID uint) error
}

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{db: db}
}

func (r *AccountRepositoryImpl) GetAccountByEmail(email string) (*model.Account, error) {
	var account model.Account
	err := r.db.Where("email = ?", email).First(&account).Error
	return &account, err
}

func (r *AccountRepositoryImpl) GetAccountById(id uint) (*model.Account, error) {
	var account model.Account
	err := r.db.First(&account, id).Error
	return &account, err
}

func (r *AccountRepositoryImpl) CreateAccount(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepositoryImpl) UpdateAccount(account *model.Account) error {
	return r.db.Save(account).Error
}

func (r *AccountRepositoryImpl) DeleteAccount(accountID uint) error {
	return r.db.Delete(&model.Account{}, accountID).Error
}
