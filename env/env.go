package env

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
