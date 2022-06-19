package persistence

import (
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

var _ repository.Account = (*AccountRepository)(nil)

func (r *AccountRepository) GetAccountByEmail(email string) (*entity.Account, error) {
	var account entity.Account
	if err := r.db.Where("email = ?", email).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) GetAccountByID(id uint) (*entity.Account, error) {
	var account entity.Account
	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) CreateAccount(account *entity.Account) error {
	if err := r.db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) UpdateAccount(account *entity.Account) error {
	if err := r.db.Save(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) DeleteAccount(accountID uint) error {
	if err := r.db.Delete(&entity.Account{}, accountID).Error; err != nil {
		return err
	}
	return nil
}
