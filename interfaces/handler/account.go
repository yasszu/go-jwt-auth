package handler

import (
	"net/http"

	"github.com/yasszu/go-jwt-auth/application/usecase"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/interfaces/middleware"
	"github.com/yasszu/go-jwt-auth/interfaces/presenter"
	"github.com/yasszu/go-jwt-auth/interfaces/response"
)

type AccountHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(accountRepository repository.AccountRepository) *AccountHandler {
	return &AccountHandler{
		accountUsecase: usecase.NewAccountUsecase(accountRepository),
	}
}

// Me  GET /v1/me
func (h *AccountHandler) Me(w http.ResponseWriter, r *http.Request) {
	accountID, ok := middleware.GetAccountID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	account, err := h.accountUsecase.Me(r.Context(), accountID)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccount(account))
}
