package model

import "gorm.io/gorm"

/*
*
post表存储博客文章信息
*/
type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title" binding:"required" json:"title"`
	Content string `gorm:"not null" json:"content" binding:"required" json:"content"`
	UserID  uint   `json:"userID"` // 外键关联users表id
	User    User   `json:"user"`
}
