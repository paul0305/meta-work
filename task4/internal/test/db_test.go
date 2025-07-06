package test

import (
	"task4/internal/config"
	"task4/internal/model"
	"testing"
)

func TestDb(t *testing.T) {
	//启动gin项目
	config.DB = config.InitDb("root", "admin123", "127.0.0.1", 3306, "blog")
	// 自动迁移模型
	err := config.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		t.Fatal(err)
	}
}
