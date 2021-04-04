package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go-jwt-auth/jwt"
	"go-jwt-auth/model"
	"go-jwt-auth/util"
)

// Signup POST /signup
func (h *Handler) Signup(c echo.Context) error {
	var form model.SignupForm
	if err := c.Bind(&form); err != nil {
		return c.String(http.StatusBadRequest, "BadRequest")
	}
	if err := c.Validate(&form); err != nil {
		return c.String(http.StatusBadRequest, "Validation Error")
	}

	var account model.Account
	if err := account.Populate(&form); err != nil {
		return c.JSON(http.StatusInternalServerError, Err(err))
	}
	if err := h.accountRepository.CreateAccount(&account); err != nil {
		return c.JSON(http.StatusInternalServerError, Err(err))
	}

	token, err := jwt.Sign(&account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err(err))
	}

	return c.JSON(http.StatusOK, token)
}

// Login POST /login
func (h *Handler) Login(c echo.Context) error {
	var form model.LoginForm
	if err := c.Bind(&form); err != nil {
		return c.String(http.StatusBadRequest, "BadRequest")
	}

	account, err := h.accountRepository.GetAccountByEmail(form.Email)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found email")
	}
	if err := util.ComparePassword(account.PasswordHash, form.Password); err != nil {
		return c.String(http.StatusForbidden, "Invalid password")
	}

	token, err := jwt.Sign(account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err(err))
	}

	return c.JSON(http.StatusOK, token)
}

// Me  GET /v1/me
func (h *Handler) Me(c echo.Context) error {
	accountID := jwt.BindUser(c).AccountID
	account, err := h.accountRepository.GetAccountById(accountID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model.NewAccountResponse(account))
}
