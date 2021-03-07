package util

import (
	"fmt"
	"go-jwt-auth/env"
)

var (
	conf Conf
)

func init() {
	conf.loadEnv()
}

// Configuration
type Conf struct {
	Server struct {
		Port string
		Host string
	}
	Postgres struct {
		Host     string
		Port     int
		Username string
		Password string
		DB       string
	}
	JWT struct {
		Secret string
	}
}

func NewConf() *Conf {
	return &conf
}

func (c *Conf) loadEnv() {
	fmt.Println("load env...")
	s := env.NewServer()
	c.Server.Host = s.Host.Value
	c.Server.Port = s.Port.Value

	pg := env.NewPostgres()
	c.Postgres.Host = pg.Host.Value
	c.Postgres.Port = pg.Port.Int()
	c.Postgres.Username = pg.User.Value
	c.Postgres.Password = pg.Password.Value
	c.Postgres.DB = pg.DB.Value

	jwt := env.NewJWT()
	c.JWT.Secret = jwt.Secret.Value
}
