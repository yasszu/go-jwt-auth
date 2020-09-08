package handler

import (
	"github.com/labstack/echo"
	"net/http"

	"go-jwt-auth/config"
	"go-jwt-auth/jwt"
	"go-jwt-auth/model"
	"go-jwt-auth/repository"
	"go-jwt-auth/util"
)

type IAccountHandler interface {
	Signup(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Verify(c echo.Context) error
}

type AccountHandler struct {
	accountRepository repository.IAccountRepository
	conf              *config.Config
}

func NewAccountHandler(repository repository.IAccountRepository, conf *config.Config) *AccountHandler {
	return &AccountHandler{accountRepository: repository, conf: conf}
}

// Signup POST /signup
func (h *AccountHandler) Signup(c echo.Context) error {
	secret := h.conf.JWT.Secret

	form := new(model.AccountForm)
	if err := c.Bind(form); err != nil {
		return c.String(http.StatusBadRequest, "BadRequest")
	}
	if err := c.Validate(form); err != nil {
		return c.String(http.StatusBadRequest, "Validation Error")
	}

	account, err := form.ToAccount()
	if err != nil {
		return err
	}

	if err := h.accountRepository.CreateAccount(account); err != nil {
		return err
	}

	token, err := jwt.Sign(form.Email, account.ID, secret)
	if err != nil {
		return err
	}

	util.SaveAuthorizationCookie(token, c)
	return c.JSON(http.StatusOK, model.NewAccountResponse(account))
}

// Login POST /login
func (h *AccountHandler) Login(c echo.Context) error {
	secret := h.conf.JWT.Secret
	email := c.FormValue("email")
	password := c.FormValue("password")

	account, err := h.accountRepository.GetAccountByEmail(email)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found email")
	}

	if err := util.ComparePassword(account.PasswordHash, password); err != nil {
		return c.String(http.StatusForbidden, "Invalid password")
	}

	token, err := jwt.Sign(account.Email, account.ID, secret)
	if err != nil {
		return err
	}

	util.SaveAuthorizationCookie(token, c)
	return c.JSON(http.StatusOK, model.NewAccountResponse(account))
}

// Logout Get /logout
func (h *AccountHandler) Logout(c echo.Context) error {
	util.DeleteAuthorizationCookie(c)
	return c.String(http.StatusOK, "Logout success")
}

// Me  Get /v1/me
func (h *AccountHandler) Me(c echo.Context) error {
	accountID := jwt.Verify(c)
	account, err := h.accountRepository.GetAccountById(accountID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.NewAccountResponse(account))
}
