package handler

import (
	"go-jwt-auth/interfaces/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) Register(e *echo.Echo) {
	// -> /
	root := e.Group("")
	root.GET("/", h.Index)
	root.GET("/healthy", h.Healthy)
	root.GET("/ready", h.Ready)
	root.POST("/signup", h.Signup)
	root.POST("/login", h.Login)

	// -> /v1/
	v1 := e.Group("/v1")
	v1.Use(middleware.HeaderAuthMiddleware())
	v1.GET("/me", h.Me)
}
