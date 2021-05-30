package handler

import (
	"go-jwt-auth/domain/entity"
	"go-jwt-auth/interfaces/response"
	"net/http"
)

// Signup POST /signup
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	form := entity.SignupForm{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	var account entity.Account
	if err := account.Populate(&form); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.accountUsecase.SignUp(r.Context(), account)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
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
		response.Error(w, http.StatusForbidden, "Invalid password")
		return
	}

	response.JSON(w, http.StatusOK, token)
}

// Me  GET /v1/me
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	accountID := 1

	account, err := h.accountUsecase.Me(r.Context(), uint(accountID))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, account)
}
