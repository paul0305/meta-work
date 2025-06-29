package service

import (
	"fmt"
	"testing"
)

func TestFindOnlyOnce(t *testing.T) {
	var a [8]int = [8]int{1, 2, 2, 4, 5, 5, 7, 8}
	findOnlyOnce(a)
}
func TestCheckIsHuiwen(t *testing.T) {
	var a int64 = 12321
	isHuiwen := checkIsHuiwen(a)
	if isHuiwen {
		fmt.Println("是回文数")
	} else {
		fmt.Println("不是回文数")
	}

}

func TestImStack(t *testing.T) {
	ImStack()
}

func TestIsUserFulStr(t *testing.T) {
	str := "()[}]{"
	isUserFulStr := IsUserFulStr(str)
	if isUserFulStr {
		fmt.Println("字符串有效")
	} else {
		fmt.Println("字符串无效")
	}
}
