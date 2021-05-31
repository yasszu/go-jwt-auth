package util

import (
	"fmt"
	"go-jwt-auth/util/env"
)

var (
	conf Conf
)

func init() {
	conf.loadEnv()
}

// Conf Configuration
type Conf struct {
	Server
	Postgres
	JWT
}

type Server struct {
	Port string
	Host string
}

func (s Server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

type JWT struct {
	Secret string
}

func NewConf() *Conf {
	return &conf
}

func (c *Conf) loadEnv() {
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
