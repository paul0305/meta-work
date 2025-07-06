package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB         *gorm.DB
	SecrectKey string = "helloJsonToken69lk"
)

func InitDb(userName string, pass string, host string, port int, dbName string) *gorm.DB {
	var err error
	var url string
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, pass, host, port, dbName)
	DB, err = gorm.Open(mysql.Open(url),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.Logger = logger.Default.LogMode(logger.Info)
	return DB
}
