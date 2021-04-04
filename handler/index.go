package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index Handler
func (h *Handler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, OK())
}

// Healthy is used for liveness probes
func (h *Handler) Healthy(c echo.Context) error {
	return c.JSON(http.StatusOK, OK())
}

// Ready is used for readiness probes
func (h *Handler) Ready(c echo.Context) error {
	var i int
	if err := h.db.Raw("SELECT 1").Scan(&i).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Err(err))
	}
	return c.JSON(http.StatusOK, OK())
}
