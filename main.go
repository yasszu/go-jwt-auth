package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-jwt-auth/conf"
	"go-jwt-auth/handler"
	"go-jwt-auth/jwt"
	"go-jwt-auth/repository"
	"go-jwt-auth/util"
)

func main() {
	cnf, err := conf.NewConf()
	if err != nil {
		panic(err.Error())
	}

	// Init Postgres
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cnf.Postgres.Host,
		cnf.Postgres.Port,
		cnf.Postgres.Username,
		cnf.Postgres.DB,
		cnf.Postgres.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	accountRepository := repository.NewAccountRepository(db)
	accountHandler := handler.NewAccountHandler(accountRepository, cnf)

	// Route => handler

	// /..
	e.GET("/", handler.Index)
	e.POST("/signup", accountHandler.Signup)
	e.POST("/login", accountHandler.Login)
	e.POST("/logout", accountHandler.Logout)

	// /v1/..
	v1 := e.Group("/v1") // Restricted group
	v1.Use(middleware.JWTWithConfig(jwt.MiddlewareConfig(cnf.JWT.Secret)))
	v1.GET("/me", accountHandler.Me)

	// Start server
	e.Logger.Fatal(e.Start(cnf.Server.Host + ":" + cnf.Server.Port))
}
