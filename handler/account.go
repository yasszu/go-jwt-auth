package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"go-jwt-auth/config"
	"go-jwt-auth/data"
	"go-jwt-auth/jwt"
	"go-jwt-auth/util"
)

func Signup(c echo.Context) error {
	secret := config.GetConfig(c).JWT.Secret
	email := c.FormValue("email")
	password := c.FormValue("password")

	accounts := data.NewAccounts(c)
	id, err := accounts.CreateAccount(email, password)
	if err != nil {
		return err
	}

	token, err := jwt.Sign(email, id, secret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// Login handler
func Login(c echo.Context) error {
	secret := config.GetConfig(c).JWT.Secret
	email := c.FormValue("email")
	password := util.Password(c.FormValue("password"))

	accounts := data.NewAccounts(c)
	a, err := accounts.GetAccountByEmail(email)
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

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// Verify handler
func Verify(c echo.Context) error {
	accountID := jwt.Verify(c)
	return c.JSON(http.StatusOK, echo.Map{
		"account_id": accountID,
	})
}
