package conf

import (
	"fmt"

	"github.com/yasszu/go-jwt-auth/util/env"
)

var (
	Server       *server
	Postgres     *postgres
	PostgresTest *postgres
	JWT          *jWT
)

func init() {
	v := env.NewVariables()
	Server = &server{
		Host: v.ServerHost.String(),
		Port: v.ServerPort.String(),
	}
	Postgres = &postgres{
		Host:     v.PostgresHost.String(),
		Port:     v.PostgresPort.Int(),
		Username: v.PostgresUser.String(),
		Password: v.PostgresPassword.String(),
		DB:       v.PostgresDB.String(),
	}
	PostgresTest = &postgres{
		Host:     v.PostgresTestHost.String(),
		Port:     v.PostgresTestPort.Int(),
		Username: v.PostgresUser.String(),
		Password: v.PostgresPassword.String(),
		DB:       v.PostgresDB.String(),
	}
	JWT = &jWT{
		Secret: v.JWTSecret.String(),
	}
}

type server struct {
	Port string
	Host string
}

func (s *server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

func (p *postgres) DSN() string {
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

func (j *jWT) SigningKey() []byte {
	return []byte(j.Secret)
}
