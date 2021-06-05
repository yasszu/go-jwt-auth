package main

import (
	"log"
	"net/http"
	"time"

	"go-jwt-auth/infrastructure/db"
	"go-jwt-auth/infrastructure/persistence"
	"go-jwt-auth/interfaces/handler"
	_middleware "go-jwt-auth/interfaces/middleware"
	"go-jwt-auth/util"

	"github.com/gorilla/mux"
)

func main() {
	// Load conf
	cnf := util.NewConf()

	// Establish DB connection
	conn, err := db.NewConn(cnf)
	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()
	middleware := _middleware.NewMiddleware()
	root := r.PathPrefix("").Subrouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.JWT)

	accountRepository := persistence.NewAccountRepository(conn)

	indexHandler := handler.NewIndexHandler(conn)
	indexHandler.Register(root)

	accountHandler := handler.NewAccountHandler(conn, accountRepository)
	accountHandler.Register(root, v1)

	srv := &http.Server{
		Handler:      r,
		Addr:         cnf.Server.Addr(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start server
	log.Println(" ⇨ http server started on", cnf.Server.Addr())
	log.Fatal(srv.ListenAndServe())
}
