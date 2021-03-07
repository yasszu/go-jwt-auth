package util

import (
	"fmt"
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
	NewServerEnv().Populate(c)
	NewPostgresEnv().Populate(c)
	NewJWTEnv().Populate(c)
}
