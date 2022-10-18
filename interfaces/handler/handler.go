package handler

import (
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/domain/service"
	"github.com/yasszu/go-jwt-auth/interfaces/middleware"
	"gorm.io/gorm"
)

type Handler struct {
	*IndexHandler
	*AccountHandler
	*AuthenticationHandler
	*middleware.Middleware
}

func NewHandler(db *gorm.DB, accountRepository repository.Account, jwtService service.Jwt) *Handler {
	indexHandler := NewIndexHandler(db)
	accountHandler := NewAccountHandler(accountRepository, jwtService)
	authenticationHandler := NewAuthenticationHandler(accountRepository, jwtService)

	return &Handler{
		IndexHandler:          indexHandler,
		AccountHandler:        accountHandler,
		AuthenticationHandler: authenticationHandler,
		Middleware:            middleware.NewHandler(jwtService),
	}
}
