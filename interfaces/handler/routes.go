package handler

import "github.com/gorilla/mux"

func (h Handler) Register(r *mux.Router) {
	root := r.PathPrefix("").Subrouter()

	root.HandleFunc("/", h.Index).Methods("GET")
	root.HandleFunc("/healthy", h.Healthy).Methods("GET")
	root.HandleFunc("/ready", h.Ready).Methods("GET")
	//root.POST("/signup", h.Signup)
	//root.POST("/login", h.Login)

	//v1 := e.Group("/v1")
	//v1.Use(middleware.HeaderAuthMiddleware())
	//v1.GET("/me", h.Me)
}
