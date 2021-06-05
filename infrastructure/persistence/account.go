package persistence

import (
	"go-jwt-auth/domain/entity"
	"go-jwt-auth/domain/repository"

	"gorm.io/gorm"
)

var _ repository.AccountRepository = &AccountRepository{}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetAccountByEmail(email string) (*entity.Account, error) {
	var account entity.Account
	err := r.db.Where("email = ?", email).First(&account).Error
	return &account, err
}

func (r *AccountRepository) GetAccountByID(id uint) (*entity.Account, error) {
	var account entity.Account
	err := r.db.First(&account, id).Error
	return &account, err
}

func (r *AccountRepository) CreateAccount(account *entity.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) UpdateAccount(account *entity.Account) error {
	return r.db.Save(account).Error
}

func (r *AccountRepository) DeleteAccount(accountID uint) error {
	return r.db.Delete(&entity.Account{}, accountID).Error
}
