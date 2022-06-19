package handler

import (
	"github.com/gorilla/mux"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/domain/service"
	"github.com/yasszu/go-jwt-auth/interfaces/middleware"
	"gorm.io/gorm"
)

type Handler struct {
	*IndexHandler
	*AccountHandler
	*AuthenticationHandler
	*middleware.Middleware
}

func NewHandler(db *gorm.DB, accountRepository repository.Account, jwtService service.Jwt) *Handler {
	indexHandler := NewIndexHandler(db)
	accountHandler := NewAccountHandler(accountRepository, jwtService)
	authenticationHandler := NewAuthenticationHandler(accountRepository, jwtService)

	return &Handler{
		IndexHandler:          indexHandler,
		AccountHandler:        accountHandler,
		AuthenticationHandler: authenticationHandler,
		Middleware:            middleware.NewHandler(jwtService),
	}
}

func (h *Handler) Register(r *mux.Router) {
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
}
