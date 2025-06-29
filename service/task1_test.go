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

func TestFindLongPubStr(t *testing.T) {
	var a [3]string = [3]string{"abcdefg", "abcd", "abcdefgh"}
	pubStr := findLongPubStr(a)
	t.Log("最长公共前缀:", pubStr)
}

// nums = [0,0,1,1,1,2,2,3,3,4]
func TestRemoveDuplicate(t *testing.T) {
	var nums []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	numLen := removeDuplicate(nums)
	t.Log("去重后数组长度:", numLen)
}

func TestFindManzu(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	manzu := findManzu(nums, 9)
	for _, v := range manzu {
		t.Log(v)
	}
}
