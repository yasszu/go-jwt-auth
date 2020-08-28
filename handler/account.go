package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

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

// Signup -> POST /signup
func (h AccountHandler) Signup(c echo.Context) error {
	secret := h.conf.JWT.Secret
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	from := model.AccountForm{Username: username, Email: email, Password: password}
	if err := c.Validate(from); err != nil {
		return err
	}

	account, err := h.accountRepository.CreateAccount(from)
	if err != nil {
		return err
	}

	token, err := jwt.Sign(email, account.AccountID, secret)
	if err != nil {
		return err
	}

	util.CookieStore{Key: "Authorization", Value: token, ExpireTime: time.Hour * 60 * 99}.Write(c)
	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// Login -> POST /login
func (h AccountHandler) Login(c echo.Context) error {
	secret := h.conf.JWT.Secret
	email := c.FormValue("email")
	password := util.Password(c.FormValue("password"))

	a, err := h.accountRepository.GetAccountByEmail(email)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found email")
	}

	if a.Password != password.SHA256() {
		return c.String(http.StatusForbidden, "Invalid password")
	}

	token, err := jwt.Sign(a.Email, a.AccountID, secret)
	if err != nil {
		return err
	}

	util.CookieStore{Key: "Authorization", Value: token, ExpireTime: time.Hour * 60 * 99}.Write(c)
	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// Logout -> Get /logout
func (h AccountHandler) Logout(c echo.Context) error {
	util.CookieStore{Key: "Authorization"}.Delete(c)
	return c.String(http.StatusOK, "Logout success")
}

// Me -> Get /v1/me
func (h AccountHandler) Me(c echo.Context) error {
	accountId := jwt.Verify(c)
	account, err := h.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	response := &model.AccountResponse{
		AccountID: account.AccountID,
		Username:  account.Username,
		Email:     account.Email,
		CreatedOn: account.CreatedOn,
	}
	return c.JSON(http.StatusOK, response)
}
