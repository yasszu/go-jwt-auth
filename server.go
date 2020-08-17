package main

import (
	"database/sql"
	"fmt"
	"go-jwt-auth/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"

	"go-jwt-auth/config"
	"go-jwt-auth/handler"
	"go-jwt-auth/jwt"
)

func main() {
	conf, err := config.NewConfig().Load()
	if err != nil {
		panic(err.Error())
	}

	// Init Postgres
	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		conf.Postgres.Username,
		conf.Postgres.DB,
		conf.Postgres.Password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	accountRepository := repository.NewAccountRepository(db)
	accountHandler := handler.NewAccountHandler(accountRepository, conf)

	// Route => handler

	// /..
	e.GET("/", handler.Index)
	e.POST("/signup", accountHandler.Signup)
	e.POST("/login", accountHandler.Login)
	e.POST("/logout", accountHandler.Logout)

	// /v1/..
	v1 := e.Group("/v1") // Restricted group
	v1.Use(middleware.JWTWithConfig(jwt.MiddlewareConfig(conf.JWT.Secret)))
	v1.GET("", handler.Index)
	v1.GET("/verify", accountHandler.Verify)

	// Start server
	e.Logger.Fatal(e.Start(conf.Server.Host + ":" + conf.Server.Port))
}
