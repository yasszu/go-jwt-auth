package db

import (
	"fmt"
	"go-jwt-auth/util"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	retryTimes  = 20
	waitingTime = 3 * time.Second
)

func NewConn(cnf *util.Conf) (*gorm.DB, error) {
	dialector := getDialector(cnf)
	config := getConfig()

	db, err := gorm.Open(dialector, config)
	for i := 0; i < retryTimes; i++ {
		if err == nil {
			break
		}
		fmt.Println("Waiting for getting the connection of Postgres...")
		time.Sleep(waitingTime)
		db, err = gorm.Open(dialector, config)
	}

	return db, err
}

func getDialector(cnf *util.Conf) gorm.Dialector {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cnf.Postgres.Host,
		cnf.Postgres.Port,
		cnf.Postgres.Username,
		cnf.Postgres.DB,
		cnf.Postgres.Password)
	return postgres.Open(dsn)
}

func getConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
}
