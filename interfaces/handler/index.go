package handler

import (
	"go-jwt-auth/interfaces/response"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type IndexHandler struct {
	db *gorm.DB
}

func NewIndexHandler(db *gorm.DB) *IndexHandler {
	return &IndexHandler{db: db}
}

func (h IndexHandler) Register(root *mux.Router) {
	root.HandleFunc("/", h.Index).Methods("GET")
	root.HandleFunc("/healthy", h.Healthy).Methods("GET")
	root.HandleFunc("/ready", h.Ready).Methods("GET")
}

// Index AccountHandler
func (h *IndexHandler) Index(w http.ResponseWriter, _ *http.Request) {
	response.OK(w)
}

// Healthy is used for liveness probes
func (h *IndexHandler) Healthy(w http.ResponseWriter, _ *http.Request) {
	response.OK(w)
}

// Ready is used for readiness probes
func (h *IndexHandler) Ready(w http.ResponseWriter, _ *http.Request) {
	var i int
	if err := h.db.Raw("SELECT 1").Scan(&i).Error; err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(w)
}
