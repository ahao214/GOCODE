package main

import (
	"fmt"
	"sort"
)

//求一个串中出现的第一个最长重复子串

//找出最长的公共子串的长度
func maxPrefix(s1, s2 string) int {
	i := 0
	for i < len(s1) && i < len(s2) {
		if s1[i] == s2[i] {
			i++
		} else {
			break
		}
	}
	return i
}

//获取最长的公共子串
func getMaxCommonStr(txt string) string {
	n := len(txt)
	//用来存储后缀数组
	suffixes := make([]string, n)
	longestSubStrLen := 0
	var longestSubStr string
	//获取到后缀数组
	for i := 0; i < n; i++ {
		suffixes[i] = txt[i:]
	}
	sort.Strings(suffixes)
	for i := 1; i < n; i++ {
		tmp := maxPrefix(suffixes[i], suffixes[i-1])
		if tmp > longestSubStrLen {
			longestSubStrLen = tmp
			longestSubStr = (suffixes[i])[0 : i+1]
		}
	}
	return longestSubStr
}

func main() {
	txt := "banana"
	fmt.Println("最长的公共子串为：", getMaxCommonStr(txt))

}
