package mysql

import (
	"fmt"

	"github.com/yasszu/go-jwt-auth/util/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConn() (*gorm.DB, error) {
	fmt.Println(conf.Mysql.DSN())
	dialector := mysql.Open(conf.Mysql.DSN())
	cnf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(dialector, cnf)
	if err != nil {
		return nil, err
	}

	return db, nil
}
