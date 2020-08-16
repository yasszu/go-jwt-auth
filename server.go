package main

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"

	"go-jwt-auth/config"
	"go-jwt-auth/handler"
	"go-jwt-auth/jwt"
)

func main() {
	conf := config.LoadConfig()
	host := conf.Server.Host
	port := conf.Server.Port

	// Init Database
	conn := "user=postgres dbname=postgres password=root sslmode=disable"
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
	e.GET("/", handler.Index)
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)

	// Configure middleware with the custom claims type
	jwtConfig := middleware.JWTConfig{
		Claims:     &jwt.CustomClaims{},
		SigningKey: []byte(conf.JWT.Secret),
		TokenLookup: "cookie:Authorization",
	}

	// Restricted group
	v1 := e.Group("/v1")
	v1.Use(middleware.JWTWithConfig(jwtConfig))

	v1.GET("", handler.Index)
	v1.GET("/verify", handler.Verify)

	// Start server
	e.Logger.Fatal(e.Start(host + ":" + port))
}
