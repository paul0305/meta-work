package task3

import "time"

type Student struct {
	Id           int64 `gorm:"primaryKey"`
	Name         string
	Email        string
	Age          int
	Birthday     time.Time
	MemberNumber string `gorm:"column:member_number"`
	Grade        string
}

func (Student) TableName() string {
	return "students"
}
