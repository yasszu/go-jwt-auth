package handler

import (
	"go-jwt-auth/interfaces/response"
	"net/http"
)

// Index Handler
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, OK())
}

// Healthy is used for liveness probes
func (h *Handler) Healthy(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, OK())
}

// Ready is used for readiness probes
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	var i int
	if err := h.db.Raw("SELECT 1").Scan(&i).Error; err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
	}
	response.JSON(w, http.StatusOK, OK())
}
