package main

import "fmt"

//392. 判断子序列
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
