package handler

import "github.com/gorilla/mux"

func (h Handler) Register(r *mux.Router) {
	root := r.PathPrefix("").Subrouter()
	root.HandleFunc("/", h.Index).Methods("GET")
	root.HandleFunc("/healthy", h.Healthy).Methods("GET")
	root.HandleFunc("/ready", h.Ready).Methods("GET")
	root.HandleFunc("/signup", h.Signup).Methods("POST")
	root.HandleFunc("/login", h.Login).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/me", h.Me).Methods("GET")
}
