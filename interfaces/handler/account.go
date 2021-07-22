package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yasszu/go-jwt-auth/application/usecase"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/infrastructure/jwt"
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
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, account.Response())
}
