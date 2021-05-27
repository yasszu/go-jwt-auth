package jwt

import (
	"go-jwt-auth/domain/entity"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	expireHour = 24 * 121
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwt.StandardClaims
}

func GetSigningKey() []byte {
	defaultKey := "b5a636fc-bd01-41b1-9780-7bbd906fa4c0"
	_ = os.Setenv("JWT_SECRET", defaultKey)
	secret := os.Getenv("JWT_SECRET")
	return []byte(secret)
}

func Sign(account *entity.Account) (*entity.AccessToken, error) {
	expiredAt := time.Now().Add(time.Hour * expireHour)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(GetSigningKey())
	if err != nil {
		return nil, err
	}
	accessToken := &entity.AccessToken{
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
