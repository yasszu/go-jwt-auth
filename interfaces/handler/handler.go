package handler

import (
	"github.com/gorilla/mux"
	"github.com/yasszu/go-jwt-auth/infrastructure/persistence"
	_middleware "github.com/yasszu/go-jwt-auth/interfaces/middleware"
	"gorm.io/gorm"
)

type Handler struct {
	*IndexHandler
	*AccountHandler
	*AuthenticationHandler

	middleware *_middleware.Middleware
}

func NewHandler(db *gorm.DB) *Handler {
	middleware := _middleware.NewMiddleware()
	accountRepository := persistence.NewAccountRepository(db)
	indexHandler := NewIndexHandler(db)
	accountHandler := NewAccountHandler(accountRepository)
	authenticationHandler := NewAuthenticationHandler(accountRepository)

	return &Handler{
		IndexHandler:          indexHandler,
		AccountHandler:        accountHandler,
		AuthenticationHandler: authenticationHandler,
		middleware:            middleware,
	}
}

func (h *Handler) Register(r *mux.Router) {
	root := r.PathPrefix("").Subrouter()
	root.Use(h.middleware.Logging)
	root.Use(h.middleware.CORS)
	root.HandleFunc("/", h.Index).Methods("GET")
	root.HandleFunc("/healthy", h.Healthy).Methods("GET")
	root.HandleFunc("/ready", h.Ready).Methods("GET")
	root.HandleFunc("/signup", h.Signup).Methods("POST")
	root.HandleFunc("/login", h.Login).Methods("POST")
	root.HandleFunc("/token", h.RefreshToken).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(h.middleware.Logging)
	v1.Use(h.middleware.JWT)
	v1.HandleFunc("/me", h.Me).Methods("GET")
}
