package router

import (
	"github.com/gorilla/mux"
	"github.com/yasszu/go-jwt-auth/interfaces/handler"
	"github.com/yasszu/go-jwt-auth/interfaces/middleware"
)

func NewRouter(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	root := r.PathPrefix("").Subrouter()
	root.Use(middleware.Logging)
	root.Use(middleware.CORS)
	root.HandleFunc("/", h.Index).Methods("GET")
	root.HandleFunc("/healthy", h.Healthy).Methods("GET")
	root.HandleFunc("/ready", h.Ready).Methods("GET")
	root.HandleFunc("/signup", h.Signup).Methods("POST")
	root.HandleFunc("/login", h.Login).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.Logging)
	v1.Use(h.JWT())
	v1.HandleFunc("/me", h.Me).Methods("GET")
	return r
}
