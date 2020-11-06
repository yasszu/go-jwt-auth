package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index Handler
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
