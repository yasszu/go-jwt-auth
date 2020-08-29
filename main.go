package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-jwt-auth/config"
	"go-jwt-auth/handler"
	"go-jwt-auth/jwt"
	"go-jwt-auth/repository"
	"go-jwt-auth/util"
)

func main() {
	conf, err := config.NewConfig().Load()
	if err != nil {
		panic(err.Error())
	}

	// Init Postgres
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.Username,
		conf.Postgres.DB,
		conf.Postgres.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Echo instance
	e := echo.New()
	e.Validator = util.NewCustomValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	accountRepository := repository.NewAccountRepository(db)
	accountHandler := handler.NewAccountHandler(accountRepository, &conf)

	// Route => handler

	// /..
	e.GET("/", handler.Index)
	e.POST("/signup", accountHandler.Signup)
	e.POST("/login", accountHandler.Login)
	e.POST("/logout", accountHandler.Logout)

	// /v1/..
	v1 := e.Group("/v1") // Restricted group
	v1.Use(middleware.JWTWithConfig(jwt.MiddlewareConfig(conf.JWT.Secret)))
	v1.GET("/me", accountHandler.Me)

	// Start server
	e.Logger.Fatal(e.Start(conf.Server.Host + ":" + conf.Server.Port))
}
