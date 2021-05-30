package main

import (
	"fmt"
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

	//// Echo instance
	//e := echo.New()
	//e.Validator = util.NewValidator()
	//
	//// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//}))

	h := handler.NewHandler(conn)
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	h.Register(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", cnf.Server.Host, cnf.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start server
	log.Fatal(srv.ListenAndServe())
}
