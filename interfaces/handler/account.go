package handler

import (
	"go-jwt-auth/domain/entity"
	"go-jwt-auth/infrastructure/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Signup POST /signup
func (h *Handler) Signup(c echo.Context) error {
	var form entity.SignupForm
	if err := c.Bind(&form); err != nil {
		return c.String(http.StatusBadRequest, "BadRequest")
	}
	if err := c.Validate(&form); err != nil {
		return c.String(http.StatusBadRequest, "Validation Error")
	}

	var account entity.Account
	if err := account.Populate(&form); err != nil {
		return c.JSON(http.StatusInternalServerError, Err(err))
	}

	token, err := h.accountUsecase.SignUp(c, account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err(err))
	}

	return c.JSON(http.StatusOK, token)
}

// Login POST /login
func (h *Handler) Login(c echo.Context) error {
	var form entity.LoginForm
	if err := c.Bind(&form); err != nil {
		return c.String(http.StatusBadRequest, "BadRequest")
	}

	token, err := h.accountUsecase.Login(c, form.Email, form.Password)
	if err != nil {
		return c.String(http.StatusForbidden, "Invalid password")
	}

	return c.JSON(http.StatusOK, token)
}

// Me  GET /v1/me
func (h *Handler) Me(c echo.Context) error {
	accountID := jwt.BindUser(c).AccountID

	account, err := h.accountUsecase.Me(c, accountID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account.Response())
}
