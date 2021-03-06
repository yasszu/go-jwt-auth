package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go-jwt-auth/jwt"
	"go-jwt-auth/model"
	"go-jwt-auth/repository"
	"go-jwt-auth/util"
)

type AccountHandler struct {
	accountRepository repository.AccountRepository
	conf              *util.Conf
}

func NewAccountHandler(repository repository.AccountRepository, conf *util.Conf) *AccountHandler {
	return &AccountHandler{accountRepository: repository, conf: conf}
}

func (h AccountHandler) RegisterRoot(e *echo.Echo) {
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
}

func (h AccountHandler) RegisterV1(v1 *echo.Group) {
	v1.GET("/me", h.Me)
}

// Signup POST /signup
func (h *AccountHandler) Signup(c echo.Context) error {
	var form model.SignupForm
	if err := c.Bind(&form); err != nil {
		return c.String(http.StatusBadRequest, "BadRequest")
	}
	if err := c.Validate(&form); err != nil {
		return c.String(http.StatusBadRequest, "Validation Error")
	}

	var account model.Account
	if err := account.Populate(&form); err != nil {
		return err
	}
	if err := h.accountRepository.CreateAccount(&account); err != nil {
		return err
	}

	token, err := jwt.Sign(&account)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}

// Login POST /login
func (h *AccountHandler) Login(c echo.Context) error {
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
		return err
	}

	return c.JSON(http.StatusOK, token)
}

// Me  GET /v1/me
func (h *AccountHandler) Me(c echo.Context) error {
	accountID := jwt.BindUser(c).AccountID
	account, err := h.accountRepository.GetAccountById(accountID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model.NewAccountResponse(account))
}
