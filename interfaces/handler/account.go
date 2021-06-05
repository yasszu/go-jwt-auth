package handler

import (
	"go-jwt-auth/infrastructure/jwt"
	"go-jwt-auth/interfaces/form"
	"go-jwt-auth/interfaces/response"
	"go-jwt-auth/interfaces/view"
	"net/http"
)

// Signup POST /signup
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
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

	token, err := h.accountUsecase.SignUp(r.Context(), account)
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, token)
}

// Login POST /login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	token, err := h.accountUsecase.Login(r.Context(), email, password)
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, token)
}

// Me  GET /v1/me
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
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
