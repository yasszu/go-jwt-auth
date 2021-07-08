package conf

import (
	"fmt"
	"go-jwt-auth/util/env"
)

var (
	Server   *server
	Postgres *postgres
	JWT      *jWT
)

func init() {
	v := env.NewVariables()
	Server = &server{
		Host: v.ServerHost.Value,
		Port: v.ServerPort.Value,
	}
	Postgres = &postgres{
		Host:     v.PostgresHost.Value,
		Port:     v.PostgresPort.Int(),
		Username: v.PostgresUser.Value,
		Password: v.PostgresPassword.Value,
		DB:       v.PostgresDB.Value,
	}
	JWT = &jWT{
		Secret: v.JWTSecret.Value,
	}
}

type server struct {
	Port string
	Host string
}

func (s server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

func (p postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		p.Host,
		p.Port,
		p.Username,
		p.DB,
		p.Password)
}

type jWT struct {
	Secret string
}

func (j jWT) SigningKey() []byte {
	return []byte(j.Secret)
}
