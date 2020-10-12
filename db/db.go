package db

import (
	"fmt"
	"go-jwt-auth/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConn(cnf *conf.Conf) (*gorm.DB, error) {
	dialector := getDialector(cnf)
	conn, err := gorm.Open(dialector, getConfig())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getDialector(cnf *conf.Conf) gorm.Dialector {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cnf.Postgres.Host,
		cnf.Postgres.Port,
		cnf.Postgres.Username,
		cnf.Postgres.DB,
		cnf.Postgres.Password,
	)
	return postgres.Open(dsn)
}

func getConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
}
