package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-jwt-auth/db"
	"go-jwt-auth/handler"
	"go-jwt-auth/jwt"
	"go-jwt-auth/repository"
	"go-jwt-auth/util"
)

func main() {
	// Load conf
	cnf := util.NewConf()

	// Echo instance
	e := echo.New()
	e.Validator = util.NewValidator()

	// Establish DB connection
	conn, err := db.NewConn(cnf)
	if err != nil {
		panic(err.Error())
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	index := handler.NewIndexHandler(conn)

	accountRepository := repository.NewAccountRepository(conn)
	account := handler.NewAccountHandler(accountRepository, cnf)

	// /..
	index.Register(e)
	account.RegisterRoot(e)

	// /v1/..
	v1 := e.Group("/v1")
	v1.Use(middleware.JWTWithConfig(jwt.HeaderAuthConfig()))
	account.RegisterV1(v1)

	// Start server
	e.Logger.Fatal(e.Start(cnf.Server.Host + ":" + cnf.Server.Port))
}
