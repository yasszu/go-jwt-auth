package handler

import (
	"gorm.io/gorm"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
	db *gorm.DB
}

func NewIndexHandler(db *gorm.DB) *IndexHandler {
	return &IndexHandler{db: db}
}

func (h *IndexHandler) Register(e *echo.Echo) {
	e.GET("/", h.Index)
	e.GET("/healthy", h.Healthy)
	e.GET("/ready", h.Ready)
}

// Index Handler
func (h *IndexHandler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, OK())
}

// Healthy is used for liveness probes
func (h *IndexHandler) Healthy(c echo.Context) error {
	return c.JSON(http.StatusOK, OK())
}

// Ready is used for readiness probes
func (h *IndexHandler) Ready(c echo.Context) error {
	var i int
	if err := h.db.Raw("SELECT 1").Scan(&i).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, OK())
}