package jwt

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtCookieAuth() echo.MiddlewareFunc {
	config := middleware.DefaultJWTConfig
	config.Claims = &CustomClaims{}
	config.SigningKey = getSigningKey()
	config.TokenLookup = "cookie:Authorization"
	return middleware.JWTWithConfig(config)
}

func JwtHeaderAuth() echo.MiddlewareFunc {
	config := middleware.DefaultJWTConfig
	config.Claims = &CustomClaims{}
	config.SigningKey = getSigningKey()
	return middleware.JWTWithConfig(config)
}