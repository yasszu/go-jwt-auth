package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/yasszu/go-jwt-auth/infrastructure/db"
	"github.com/yasszu/go-jwt-auth/interfaces/handler"
	"github.com/yasszu/go-jwt-auth/util/conf"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*30, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Establish DB connection
	conn, err := db.NewConn()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	h := handler.NewHandler(conn)
	h.Register(r)

	srv := &http.Server{
		Addr:         conf.Server.Addr(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Infof(" ⇨ http server started on %s", conf.Server.Addr())
		log.Infof(" ⇨ graceful timeout: %s", wait)
		if err = srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)`
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c
	log.Info("received stop signal")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer func() {
		log.Info("cancel")
		cancel()
	}()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("shutting down")
}
