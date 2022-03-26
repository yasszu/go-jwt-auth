package handler

import (
	"net/http"

	"github.com/yasszu/go-jwt-auth/interfaces/presenter"

	"gorm.io/gorm"
)

type IndexHandler struct {
	db *gorm.DB
}

func NewIndexHandler(db *gorm.DB) *IndexHandler {
	return &IndexHandler{db: db}
}

// Index AccountHandler
func (h *IndexHandler) Index(w http.ResponseWriter, _ *http.Request) {
	presenter.OK(w)
}

// Healthy is used for liveness probes
func (h *IndexHandler) Healthy(w http.ResponseWriter, _ *http.Request) {
	presenter.OK(w)
}

// Ready is used for readiness probes
func (h *IndexHandler) Ready(w http.ResponseWriter, _ *http.Request) {
	var i int
	if err := h.db.Raw("SELECT 1").Scan(&i).Error; err != nil {
		presenter.Error(w, http.StatusInternalServerError, err)
		return
	}
	presenter.OK(w)
}
