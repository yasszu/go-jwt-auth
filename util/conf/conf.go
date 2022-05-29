package conf

import (
	"github.com/yasszu/go-jwt-auth/util/env"
)

var (
	Server *server
	Mysql  *mysql
	JWT    *jWT
)

func init() {
	v := env.NewVariables()

	Server = &server{
		Host: v.ServerHost.String(),
		Port: v.ServerPort.String(),
	}

	Mysql = &mysql{
		Host:     v.MysqlHost.String(),
		Port:     v.MysqlPort.Int(),
		Username: v.MysqlUser.String(),
		Password: v.MysqlPassword.String(),
		DB:       v.MysqlDB.String(),
	}

	JWT = &jWT{
		Secret: v.JWTSecret.String(),
	}
}
