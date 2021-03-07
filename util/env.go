package util

import (
	"os"
	"strconv"
)

type Env struct {
	Key   string
	Value string
}

func (env *Env) Load() {
	if v, ok := os.LookupEnv(env.Key); ok {
		env.Value = v
	}
}

func (env Env) Int() int {
	i, err := strconv.Atoi(env.Value)
	if err != nil {
		return 0
	}
	return i
}

type ServerEnv struct {
	Host Env
	Port Env
}

func NewServerEnv() ServerEnv {
	e := ServerEnv{
		Host: Env{
			Key:   "SERVER_HOST",
			Value: "localhost",
		},
		Port: Env{
			Key:   "SERVER_PORT",
			Value: "8888",
		},
	}
	e.load()
	//loadEnvs(e)

	return e
}

func (e *ServerEnv) load() {
	e.Port.Load()
	e.Host.Load()
}

func (e ServerEnv) Populate(c *Conf) {
	c.Server.Host = e.Host.Value
	c.Server.Port = e.Port.Value
}

type PostgresEnv struct {
	Host     Env
	Port     Env
	User     Env
	Password Env
	DB       Env
}

func NewPostgresEnv() PostgresEnv {
	e := PostgresEnv{
		Host: Env{
			Key:   "POSTGRES_HOST",
			Value: "localhost",
		},
		Port: Env{
			Key:   "POSTGRES_PORT",
			Value: "5432",
		},
		User: Env{
			Key:   "POSTGRES_USER",
			Value: "postgres",
		},
		Password: Env{
			Key:   "POSTGRES_PASSWORD",
			Value: "password",
		},
		DB: Env{
			Key:   "POSTGRES_DB",
			Value: "postgres",
		},
	}
	e.load()
	return e
}

func (e *PostgresEnv) load() {
	e.Host.Load()
	e.Port.Load()
	e.Password.Load()
	e.User.Load()
	e.DB.Load()
}

func (e PostgresEnv) Populate(c *Conf) {
	c.Postgres.Host = e.Host.Value
	c.Postgres.Port = e.Port.Int()
	c.Postgres.Username = e.User.Value
	c.Postgres.Password = e.Password.Value
	c.Postgres.DB = e.DB.Value
}

type JWTEnv struct {
	Secret Env
}

func NewJWTEnv() JWTEnv {
	e := JWTEnv{Secret: Env{
		Key:   "JWT_SECRET",
		Value: "your_secret_key",
	}}
	e.load()
	return e
}

func (e *JWTEnv) load() {
	e.Secret.Load()
}

func (e JWTEnv) Populate(c *Conf) {
	c.JWT.Secret = e.Secret.Value
}
