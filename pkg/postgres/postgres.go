package postgres

import (
	"github.com/yasszu/go-jwt-auth/pkg/conf"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConn() (*gorm.DB, error) {
	dsn := conf.Postgres.DSN()
	dialector := pg.Open(dsn)
	cnf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}
	return gorm.Open(dialector, cnf)
}

func NewTestConn() (*gorm.DB, error) {
	dsn := conf.PostgresTest.DSN()
	dialector := pg.Open(dsn)
	cnf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	return gorm.Open(dialector, cnf)
}
