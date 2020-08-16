package jwt

import (
	"github.com/labstack/echo/middleware"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// CustomClaims are custom claims extending default ones.
type CustomClaims struct {
	Email     string `json:"email"`
	AccountID int64  `json:"account_id"`
	jwt.StandardClaims
}

func Sign(email string, id int64, secret string) (string, error) {
	// Set custom claims
	claims := &CustomClaims{
		email,
		id,
		jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Hour * 64 * 100).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(secret))
}

func Verify(c echo.Context) int64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*CustomClaims)
	return claims.AccountID
}

// MiddlewareConfig Configure middleware with the custom claims type
func MiddlewareConfig(secret string) middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:      &CustomClaims{},
		SigningKey:  []byte(secret),
		TokenLookup: "cookie:Authorization",
	}
}