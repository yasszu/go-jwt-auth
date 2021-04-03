package jwt

import (
	"go-jwt-auth/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwt.StandardClaims
}

const (
	expireHour = 24 * 121
)

func getSigningKey() []byte {
	defaultKey := "b5a636fc-bd01-41b1-9780-7bbd906fa4c0"
	os.Setenv("JWT_SECRET", defaultKey)
	secret := os.Getenv("JWT_SECRET")
	return []byte(secret)
}

func Sign(account *model.Account) (*model.AccessToken, error) {
	expiredAt := time.Now().Add(time.Hour * expireHour)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(getSigningKey())
	if err != nil {
		return nil, err
	}
	accessToken := &model.AccessToken{
		AccountID: account.ID,
		Token:     signedString,
		ExpiresAt: expiredAt,
	}
	return accessToken, nil
}

func BindUser(c echo.Context) *CustomClaims {
	token := c.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token)
	claims := token.Claims.(*CustomClaims)
	return claims
}
