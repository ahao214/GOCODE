package main

import "fmt"

//392. 判断子序列
/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

*/
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	first := 0
	second := 0
	for first < len(s) && second < len(t) {
		for second < len(t) && s[first] != t[second] {
			//s中的字母不等于t中的字母
			second++
		}
		if first < len(s) && second < len(t) && s[first] == t[second] {
			//相等的时候
			first++
			second++
		}
	}
	if first < len(s) && second == len(t) {
		//s中的字母没在t中
		return false
	}
	return first == len(s)
}

func main() {
	s := "abc"
	t := "adebxc"
	fmt.Println(isSubsequence(s, t))
}
