package task3

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	var stu Student
	err := stu.addStu(Student{Name: "张三", Age: 20,
		Birthday:     time.Now(),
		Grade:        "三年级",
		Email:        "paul@163.com",
		MemberNumber: "123456",
	})
	t.Log(err)
}

func TestFindStu(t *testing.T) {
	var stu Student
	stu.findStuByAge(18)

}

func TestUdateStudent(t *testing.T) {
	var stu Student
	n, err := stu.udateStudent()
	t.Log(n, err)
}

func TestDel(t *testing.T) {
	var stu Student
	n, err := stu.delStudent()
	t.Log(n, err)
}
