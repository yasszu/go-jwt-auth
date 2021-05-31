package main

import (
	"go-jwt-auth/infrastructure/db"
	"go-jwt-auth/interfaces/handler"
	"go-jwt-auth/interfaces/middleware"
	"go-jwt-auth/util"
	"log"
	"net/http"
	"time"

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

	h := handler.NewHandler(conn)
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	h.Register(r)

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
