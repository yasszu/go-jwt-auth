package conf

import "fmt"

type server struct {
	Port string
	Host string
}

func (s *server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
