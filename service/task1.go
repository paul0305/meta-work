package service

import (
	"fmt"
	"strconv"
)

/**
*给定一个非空整数数组，除了某个元素只出现一次以外，
其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，
然后再遍历 map 找到出现次数为1的元素
*/

func findOnlyOnce(arr [8]int) {
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if m[arr[i]] == 0 {
			m[arr[i]] = 1 // 记录元素出现的次数
		} else {
			m[arr[i]] = m[arr[i]] + 1 // 更新元素出现的次数
		}
	}
	//查找出现一次的元素
	fmt.Println("出现次数为1次的元素个数为:")
	for k, v := range m {
		if v == 1 {
			fmt.Printf("%d\n", k)
		}
	}
}

// 判断一个整数是否是回文数
func checkIsHuiwen(num int64) bool {
	// 将整数转为字符串
	str := strconv.FormatInt(num, 10)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			fmt.Printf("%d\n不是回文数", num)
			return false
		}
	}
	fmt.Printf("%d\n是回文数", num)
	return true

}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
/**
有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
示例 1：

输入：s = "()"

输出：true

示例 2：

输入：s = "()[]{}"

输出：true

示例 3：

输入：s = "(]"

输出：false

示例 4：

输入：s = "([])"

输出：true
*/
func IsUserFulStr(str string) bool {
	//将字符串转换成字符数组
	strArr := []rune(str)
	//创建一个栈，保存数据
	var strStack []rune
	strStack = make([]rune, 0)
	//遍历字符数组
	for _, v := range strArr {
		if v == '(' || v == '[' || v == '{' {
			//放入栈
			strStack = append(strStack, v)
		} else {
			if len(strStack) == 0 {
				return false
			}
			//获取栈顶元素
			top := strStack[len(strStack)-1]
			if (v == ')' && top == '(') || (v == ']' && top == '[') || (v == '}' && top == '{') {
				//出栈
				strStack = strStack[:len(strStack)-1]
			}

		}
	}
	if len(strStack) != 0 {
		return false
	}
	return true
}
func ImStack() {
	stack := []int{}
	stack = append(stack, 1) // Push
	stack = append(stack, 2)
	stack = append(stack, 25)
	stack = append(stack, 98)
	v := stack[len(stack)-1]
	fmt.Println("获取栈顶的元素:", v)
	stack = stack[:len(stack)-1] // Pop
	fmt.Println("After pop:", stack)
}
