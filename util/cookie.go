package util

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type CookieStore struct {
	Key        string
	Value      string
	ExpireTime time.Duration
}

func (s CookieStore) Write(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = s.Key
	cookie.Value = s.Value
	cookie.Expires = time.Now().Add(s.ExpireTime)
	c.SetCookie(cookie)
}

func (s CookieStore) Delete(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = s.Key
	cookie.Value = s.Value
	cookie.Expires = time.Now().Add(0)
	c.SetCookie(cookie)
}

const (
	authKey    = "Authorization"
	expireHour = 24 * 120
)

func SaveAuthorizationCookie(token string, c echo.Context) {
	cookie := CookieStore{Key: authKey, Value: token, ExpireTime: time.Hour * expireHour}
	cookie.Write(c)
}

func DeleteAuthorizationCookie(c echo.Context) {
	CookieStore{Key: authKey}.Delete(c)
}
