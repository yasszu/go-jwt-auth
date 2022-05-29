package conf

import "fmt"

type mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

func (m *mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}
