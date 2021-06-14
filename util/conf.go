package util

import (
	"fmt"
	"go-jwt-auth/util/env"
)

var (
	conf Conf
)

func init() {
	conf.load()
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

func (c *Conf) load() {
	v := env.NewVariables()
	c.Server.Host = v.ServerHost.Value
	c.Server.Port = v.ServerPort.Value
	c.Postgres.Host = v.PostgresHost.Value
	c.Postgres.Port = v.PostgresPort.Int()
	c.Postgres.Username = v.PostgresUser.Value
	c.Postgres.Password = v.PostgresPassword.Value
	c.Postgres.DB = v.PostgresDB.Value
	c.JWT.Secret = v.JWTSecret.Value
}
