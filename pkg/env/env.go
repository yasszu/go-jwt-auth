package env

import (
	"strconv"
)

type Env string

func (e Env) String() string {
	return string(e)
}

func (e Env) Int() int {
	i, err := strconv.Atoi(string(e))
	if err != nil {
		return 0
	}
	return i
}
