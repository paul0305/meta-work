package model

import "gorm.io/gorm"

/*
*
存储评论信息
*/
type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"userID"` //外键 user_id关联user表的id
	User    User   `json:"user"`
	PostID  uint   `json:"postID"` //外键 post_id关联post表的id
	Post    Post   `json:"post"`
}
