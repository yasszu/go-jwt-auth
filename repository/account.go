package repository

import (
	"go-jwt-auth/model"
	"go-jwt-auth/util"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountByEmail(email string) (*model.Account, error)
	GetAccountById(id int64) (*model.Account, error)
	CreateAccount(form model.AccountForm) (*model.Account, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (a AccountRepository) GetAccountByEmail(email string) (*model.Account, error) {
	var account model.Account
	err := a.db.Where("email = ?", email).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}

func (a AccountRepository) GetAccountById(id int64) (*model.Account, error) {
	var account model.Account
	err := a.db.Where("account_id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}

func (a AccountRepository) CreateAccount(form model.AccountForm) (*model.Account, error) {
	hash := util.Password(form.Password).SHA256()
	account := model.Account{Username: form.Username, Email: form.Email, Password: hash}
	err := a.db.Create(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}
