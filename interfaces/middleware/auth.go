package middleware

import (
	"go-jwt-auth/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CookieAuthMiddleware() echo.MiddlewareFunc {
	config := middleware.DefaultJWTConfig
	config.Claims = &jwt.CustomClaims{}
	config.SigningKey = jwt.GetSigningKey()
	config.TokenLookup = "cookie:Authorization"
	return middleware.JWTWithConfig(config)
}

func HeaderAuthMiddleware() echo.MiddlewareFunc {
	config := middleware.DefaultJWTConfig
	config.Claims = &jwt.CustomClaims{}
	config.SigningKey = jwt.GetSigningKey()
	return middleware.JWTWithConfig(config)
}
