package main

import (
	"context"
	"flag"
	"go-jwt-auth/util/conf"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-jwt-auth/infrastructure/db"
	"go-jwt-auth/infrastructure/persistence"
	"go-jwt-auth/interfaces/handler"
	_middleware "go-jwt-auth/interfaces/middleware"

	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*30, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Establish DB connection
	conn, err := db.NewConn()
	if err != nil {
		panic(err)
	}

	middleware := _middleware.NewMiddleware()
	accountRepository := persistence.NewAccountRepository(conn)
	indexHandler := handler.NewIndexHandler(conn)
	accountHandler := handler.NewAccountHandler(conn, accountRepository)
	authenticationHandler := handler.NewAuthenticationHandler(conn, accountRepository)

	r := mux.NewRouter()
	r.Use(middleware.CORS)
	r.Use(middleware.Logging)

	root := r.PathPrefix("").Subrouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.JWT)

	indexHandler.Register(root)
	authenticationHandler.Register(root)
	accountHandler.Register(v1)

	srv := &http.Server{
		Addr:         conf.Server.Addr(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf(" ⇨ http server started on %s", conf.Server.Addr())
		log.Printf(" ⇨ graceful timeout: %s", wait)
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)`
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c
	log.Println("received stop signal")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer func() {
		log.Println("cancel")
		cancel()
	}()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
}
