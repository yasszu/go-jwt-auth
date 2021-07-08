package db

import (
	"fmt"
	"go-jwt-auth/util/conf"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	retryTimes  = 20
	waitingTime = 3 * time.Second
)

func NewConn() (*gorm.DB, error) {
	db, err := openDB(conf.Postgres.DSN())
	for i := 0; i < retryTimes; i++ {
		if err == nil {
			break
		}
		fmt.Println("Waiting for getting the connection of postgres...")
		time.Sleep(waitingTime)
		db, err = openDB(conf.Postgres.DSN())
	}

	return db, err
}

func openDB(dsn string) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)
	cnf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	return gorm.Open(dialector, cnf)
}
