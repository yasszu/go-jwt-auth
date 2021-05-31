package handler

import (
	"go-jwt-auth/application/usecase"
	"go-jwt-auth/infrastructure/persistence"

	"gorm.io/gorm"
)

type Handler struct {
	db             *gorm.DB
	accountUsecase usecase.AccountUsecase
}

func NewHandler(db *gorm.DB) Handler {
	accountRepository := persistence.NewAccountRepository(db)
	return Handler{
		db:             db,
		accountUsecase: usecase.NewAccountUsecase(accountRepository),
	}
}
