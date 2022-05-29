package env

import (
	"os"
	"reflect"
)

type Variables struct {
	MysqlHost     Env `env:"MYSQL_HOST"`
	MysqlPort     Env `env:"MYSQL_PORT"`
	MysqlUser     Env `env:"MYSQL_USER"`
	MysqlPassword Env `env:"MYSQL_PASSWORD"`
	MysqlDB       Env `env:"MYSQL_DB"`
	ServerHost    Env `env:"SERVER_HOST"`
	ServerPort    Env `env:"SERVER_PORT"`
	JWTSecret     Env `env:"JWT_SECRET"`
}

func NewVariables() Variables {
	v := Variables{}
	v.load()
	return v
}

func (v *Variables) load() {
	value := reflect.ValueOf(v).Elem()
	valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		typeField := valueType.Field(i)
		tag := typeField.Tag
		key := tag.Get("env")
		if env, ok := os.LookupEnv(key); ok {
			field.SetString(env)
		}
	}
}
