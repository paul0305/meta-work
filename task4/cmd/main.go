package main

import (
	"task4/internal/api"
	"task4/internal/config"
)

func main() {
	//启动gin项目
	config.DB = config.InitDb("root", "admin123", "127.0.0.1", 3306, "blog")
	api.InitGin()
}
