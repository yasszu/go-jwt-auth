package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index Handler
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
