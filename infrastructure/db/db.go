package db

import (
	"github.com/yasszu/go-jwt-auth/util/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConn() (*gorm.DB, error) {
	return openDB(conf.Postgres.DSN())
}

func openDB(dsn string) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)
	cnf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	return gorm.Open(dialector, cnf)
}
