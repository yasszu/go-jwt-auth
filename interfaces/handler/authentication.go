package handler

import (
	"net/http"

	"github.com/yasszu/go-jwt-auth/application/usecase"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/interfaces/form"
	"github.com/yasszu/go-jwt-auth/interfaces/presenter"
	"github.com/yasszu/go-jwt-auth/interfaces/response"
)

type AuthenticationHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAuthenticationHandler(accountRepository repository.AccountRepository) *AuthenticationHandler {
	return &AuthenticationHandler{
		accountUsecase: usecase.NewAccountUsecase(accountRepository),
	}
}

// Signup POST /signup
func (h *AuthenticationHandler) Signup(w http.ResponseWriter, r *http.Request) {
	f := form.Signup{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if err := f.Validate(); err != nil {
		presenter.Error(w, presenter.Status(err), err)
		return
	}

	account, err := f.Entity()
	if err != nil {
		presenter.Error(w, presenter.Status(err), err)
		return
	}

	token, err := h.accountUsecase.SignUp(r.Context(), &account)
	if err != nil {
		presenter.Error(w, presenter.Status(err), err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}

// Login POST /login
func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	token, err := h.accountUsecase.Login(r.Context(), email, password)
	if err != nil {
		presenter.Error(w, presenter.Status(err), err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}
