package handler

import (
	"net/http"

	"go-jwt-auth/application/usecase"
	"go-jwt-auth/domain/repository"
	"go-jwt-auth/interfaces/form"
	"go-jwt-auth/interfaces/response"

	"github.com/gorilla/mux"
)

type AuthenticationHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAuthenticationHandler(accountRepository repository.AccountRepository) *AuthenticationHandler {
	return &AuthenticationHandler{
		accountUsecase: usecase.NewAccountUsecase(accountRepository),
	}
}

func (h *AuthenticationHandler) Register(r *mux.Router) {
	r.HandleFunc("/signup", h.Signup).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}

// Signup POST /signup
func (h *AuthenticationHandler) Signup(w http.ResponseWriter, r *http.Request) {
	f := form.Signup{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if err := f.Validate(); err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	account, err := f.Entity()
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	token, err := h.accountUsecase.SignUp(r.Context(), &account)
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, token)
}

// Login POST /login
func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	token, err := h.accountUsecase.Login(r.Context(), email, password)
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, token)
}
