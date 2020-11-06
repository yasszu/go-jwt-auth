package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CustomClaims are custom claims extending default ones.
type CustomClaims struct {
	Email     string `json:"email"`
	AccountID uint   `json:"account_id"`
	jwt.StandardClaims
}

const (
	expireHour = 24 * 121
)

func Sign(email string, id uint, secret string) (string, error) {
	expiredAt := time.Now().Add(time.Hour * expireHour).Unix()
	claims := &CustomClaims{
		Email:          email,
		AccountID:      id,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiredAt},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func BindUser(c echo.Context) *CustomClaims {
	user := c.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token)
	claims := user.Claims.(*CustomClaims)
	return claims
}

// CookieAuthConfig Configure middleware with the custom claims type
func CookieAuthConfig() middleware.JWTConfig {
	config := middleware.DefaultJWTConfig
	config.Claims = &CustomClaims{}
	config.SigningKey = getSigningKey()
	config.TokenLookup = "cookie:Authorization"
	return config
}

func getSigningKey() []byte {
	os.Setenv("JWT_SECRET", "b5a636fc-bd01-41b1-9780-7bbd906fa4c0") // Set dummy key
	secret := os.Getenv("JWT_SECRET")
	return []byte(secret)
}
