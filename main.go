package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"go-jwt-auth/conf"
	"go-jwt-auth/db"
	"go-jwt-auth/handler"
	"go-jwt-auth/jwt"
	"go-jwt-auth/repository"
	"go-jwt-auth/util"
)

func main() {
	// Load conf
	cnf, err := conf.NewConf()
	if err != nil {
		panic(err.Error())
	}

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

	accountRepository := repository.NewAccountRepository(conn)
	accountHandler := handler.NewAccountHandler(accountRepository, cnf)

	// /..
	e.GET("/", handler.Index)
	e.POST("/signup", accountHandler.Signup)
	e.POST("/login", accountHandler.Login)
	e.POST("/logout", accountHandler.Logout)

	// /v1/..
	v1 := e.Group("/v1")
	v1.Use(middleware.JWTWithConfig(jwt.MiddlewareConfig(cnf.JWT.Secret)))
	v1.GET("/me", accountHandler.Me)

	// Start server
	e.Logger.Fatal(e.Start(cnf.Server.Host + ":" + cnf.Server.Port))
}
