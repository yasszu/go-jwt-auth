package main

import (
	"database/sql"
	"fmt"

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

	pgUser := conf.Database.Username
	pgPass := conf.Database.Password
	pgDB := conf.Database.DB

	// Init Database
	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", pgUser, pgDB, pgPass)
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

	// Use config
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("conf", conf)
			return next(c)
		}
	})

	// Use DB
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// Route => handler

	// /..
	e.GET("/", handler.Index)
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)

	// /v1/..
	v1 := e.Group("/v1") // Restricted group
	v1.Use(middleware.JWTWithConfig(jwt.MiddlewareConfig(conf.JWT.Secret)))
	v1.GET("", handler.Index)
	v1.GET("/verify", handler.Verify)

	// Start server
	e.Logger.Fatal(e.Start(conf.Server.Host + ":" + conf.Server.Port))
}
