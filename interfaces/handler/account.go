package handler

import (
	"net/http"

	"go-jwt-auth/application/usecase"
	"go-jwt-auth/domain/repository"
	"go-jwt-auth/infrastructure/jwt"
	"go-jwt-auth/interfaces/response"
	"go-jwt-auth/interfaces/view"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AccountHandler struct {
	db             *gorm.DB
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(db *gorm.DB, accountRepository repository.AccountRepository) *AccountHandler {
	return &AccountHandler{
		db:             db,
		accountUsecase: usecase.NewAccountUsecase(accountRepository),
	}
}

func (h *AccountHandler) Register(r *mux.Router) {
	r.HandleFunc("/me", h.Me).Methods("GET")
}

// Me  GET /v1/me
func (h *AccountHandler) Me(w http.ResponseWriter, r *http.Request) {
	accountID, ok := jwt.GetAccountID(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	account, err := h.accountUsecase.Me(r.Context(), accountID)
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, view.NewAccount(account))
}
