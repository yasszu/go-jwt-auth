package handler

import (
	"go-jwt-auth/repository"
	"gorm.io/gorm"
)

type Handler struct {
	db                *gorm.DB
	accountRepository repository.AccountRepository
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{
		db:                db,
		accountRepository: repository.NewAccountRepository(db),
	}
}

func OK() map[string]interface{} {
	return map[string]interface{}{
		"message": "OK",
	}
}

func Err(err error) map[string]interface{} {
	return map[string]interface{}{
		"message": err.Error(),
	}
}
