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

/**
查找字符串数组中的最长公共前缀
*/

func findLongPubStr(strs [3]string) string {

	// 取第一个字符串作为基准
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		// 比较当前字符串与基准字符串的前缀
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		// 更新公共前缀
		prefix = prefix[:j]
		// 如果前缀为空，则没有公共前缀
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

/**
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，
从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
*/

func plusOne(digits []int) []int {
	// 从最低位开始，从右到左遍历
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果当前位不是9，则加1并返回结果
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		// 如果当前位是9，则将当前位改为0
		digits[i] = 0
	}
	// 如果所有位都是9，则将数组长度加1，并返回结果
	digits = append([]int{1}, digits...)
	return digits
}

/**
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次
，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数
*/

func removeDuplicate(nums []int) int {
	//定义快慢指针，慢指针指向不重复的数组位置，快指针遍历数组元素
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast] //赋值值，覆盖掉重复的元素
		}
		fast++
	}
	return slow + 1
}

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/

func findManzu(nums []int, target int) []int {
	//遍历数组，和map存储元素，判断map中是否存在target-nums[i]
	m := make(map[int]int) //创建mapkey 为数组元素，value为数组元素索引
	res := make([]int, 0, 2)
	for i, v := range nums {
		//如果遍历数组的元素和map中的元素存在，则返回索引
		if index, ok := m[target-v]; ok {
			res = append(res, index, i)
			return res
		} else {
			m[v] = i
		}
	}
	return res
}
