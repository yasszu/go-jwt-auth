package persistence

import (
	"context"

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

func (r *AccountRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account entity.Account
	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&account).
		Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, id uint) (*entity.Account, error) {
	var account entity.Account
	if err := r.db.WithContext(ctx).First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *entity.Account) error {
	if err := r.db.WithContext(ctx).Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) UpdateAccount(ctx context.Context, account *entity.Account) error {
	if err := r.db.WithContext(ctx).Save(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) DeleteAccount(ctx context.Context, accountID uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Account{}, accountID).Error; err != nil {
		return err
	}
	return nil
}
