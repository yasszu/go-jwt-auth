package router

import (
	"github.com/gorilla/mux"
	"github.com/yasszu/go-jwt-auth/interfaces/handler"
	"github.com/yasszu/go-jwt-auth/interfaces/middleware"
)

func NewRouter(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)

	r.HandleFunc("/", h.Index).Methods("GET")
	r.HandleFunc("/healthy", h.Healthy).Methods("GET")
	r.HandleFunc("/ready", h.Ready).Methods("GET")
	r.HandleFunc("/signup", h.Signup).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(h.JWT())
	v1.HandleFunc("/me", h.Me).Methods("GET")
	return r
}
