package handler

import (
	"net/http"

	"github.com/yasszu/go-jwt-auth/application/usecase"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/domain/service"
	"github.com/yasszu/go-jwt-auth/interfaces/form"
	"github.com/yasszu/go-jwt-auth/interfaces/presenter"
	"github.com/yasszu/go-jwt-auth/interfaces/response"
)

type AuthenticationHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAuthenticationHandler(accountRepository repository.Account, jwtService service.Jwt) *AuthenticationHandler {
	return &AuthenticationHandler{
		accountUsecase: usecase.NewAccountUsecase(accountRepository, jwtService),
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
		presenter.NewBadRequest(w)
		return
	}

	account, err := f.Entity()
	if err != nil {
		presenter.NewBadRequest(w)
		return
	}

	token, err := h.accountUsecase.SignUp(r.Context(), &account)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}

// Login POST /login
func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	f := form.Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	token, err := h.accountUsecase.Login(r.Context(), f.Email, f.Password)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewAccessToken(token))
}
