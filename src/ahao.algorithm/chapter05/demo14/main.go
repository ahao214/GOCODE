package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//统计字符串中连续重复字符的个数

func getMaxDupChar(s string, startIndex, curMaxLen, maxLen int) int {
	//字符串遍历结束，返回最长连续重复字符串的长度
	if startIndex == len(s)-1 {
		return Max(curMaxLen, maxLen)
	}
	//如果两个连续的字符相等，则在递归调用的时候把当前最长串的长度加1
	if s[startIndex] == s[startIndex+1] {
		return getMaxDupChar(s, startIndex+1, curMaxLen, maxLen)
	} else {
		//两个连续的子串不相等，求出最长串(max(curMaxLen,maxLen))
		//当前连续重复字符串的长度变为1
		return getMaxDupChar(s, startIndex+1, Max(curMaxLen, maxLen))
	}
}

func main() {
	fmt.Println("abbc的最长连续重复子串长度为：", getMaxDupChar("abbc", 0, 0, 1))
	fmt.Println("aaabbcc的最长连续重复子串长度为：", getMaxDupChar("aaabbcc", 0, 0, 1))

}
