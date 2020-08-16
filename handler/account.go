package handler

import (
	"go-jwt-auth/config"
	"go-jwt-auth/data"
	"go-jwt-auth/jwt"
	"go-jwt-auth/util"
	"net/http"
	"time"

	"github.com/labstack/echo"
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

	util.CookieStore{
		Key:        "Authorization",
		Value:      token,
		ExpireTime: time.Hour * 60 * 100,
	}.Write(c)

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

	util.CookieStore{
		Key:        "Authorization",
		Value:      token,
		ExpireTime: time.Hour * 60 * 100,
	}.Write(c)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func Logout(c echo.Context) error {
	util.CookieStore{Key: "Authorization"}.Delete(c)
	return c.String(http.StatusOK, "Logout success")
}

// Verify handler
func Verify(c echo.Context) error {
	accountID := jwt.Verify(c)
	return c.JSON(http.StatusOK, echo.Map{
		"account_id": accountID,
	})
}
