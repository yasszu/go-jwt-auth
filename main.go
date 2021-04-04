package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-jwt-auth/db"
	"go-jwt-auth/handler"
	"go-jwt-auth/util"
)

func main() {
	// Load conf
	cnf := util.NewConf()

	// Establish DB connection
	conn, err := db.NewConn(cnf)
	if err != nil {
		panic(err.Error())
	}

	// Echo instance
	e := echo.New()
	e.Validator = util.NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	h := handler.NewHandler(conn)
	h.Register(e)

	// Start server
	e.Logger.Fatal(e.Start(cnf.Server.Host + ":" + cnf.Server.Port))
}
