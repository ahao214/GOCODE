package main

import (
	"fmt"
	"strconv"
)

//题目：9.回文数
//说明：给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false
//输入：x = 121
//输出：true
//解释：
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	num := -121
	res := isPalindrome(num)
	fmt.Println(res)
}
