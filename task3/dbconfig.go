package task3

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb() {
	var err error
	db, err = gorm.Open(mysql.Open("root:admin123@tcp(127.0.0.1:3306)/edu_user?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
