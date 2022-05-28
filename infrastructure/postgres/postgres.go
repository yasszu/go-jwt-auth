package postgres

import (
	"github.com/yasszu/go-jwt-auth/util/conf"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConn() (*gorm.DB, error) {
	dns := conf.Postgres.DSN()
	return openDB(dns)
}

func openDB(dsn string) (*gorm.DB, error) {
	dialector := pg.Open(dsn)
	cnf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	return gorm.Open(dialector, cnf)
}
