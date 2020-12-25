package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	conf "knowledge/config"
	logging "logger"
)

var database *gorm.DB

func Connect() *gorm.DB {
	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", conf.Config.DBConf.Host, conf.Config.DBConf.Port, conf.Config.DBConf.Login, conf.Config.DBConf.Password, conf.Config.DBConf.DatabaseName)
	var err error
	database, err = gorm.Open(conf.Config.DBConf.DatabaseType, connStr)
	if err != nil {
		logging.DatabaseLog.Critical("Panic! Can't connect to database.")
		panic(err)
	}
	return database
}
