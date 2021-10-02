package handler

import (
	"net/http"

	"github.com/yasszu/go-jwt-auth/application/usecase"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/interfaces/form"
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
		response.Error(w, response.Status(err), err)
		return
	}

	account, err := f.Entity()
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	tokenPair, err := h.accountUsecase.SignUp(r.Context(), &account)
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, response.NewTokenResponse(tokenPair))
}

// Login POST /login
func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	tokenPair, err := h.accountUsecase.Login(r.Context(), email, password)
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, response.NewTokenResponse(tokenPair))
}

func (h *AuthenticationHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.FormValue("refresh_token")

	tokenPair, err := h.accountUsecase.RefreshToken(r.Context(), refreshToken)
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, response.NewTokenResponse(tokenPair))
}
