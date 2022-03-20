package main

import "fmt"

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
*/
func isAnagram(s string, t string) bool {
	hash := map[rune]int{}
	for _, value := range s {
		hash[value]++
	}
	for _, value := range t {
		hash[value]--
	}
	for _, value := range hash {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	isTrue := isAnagram("rat", "tar")
	fmt.Println(isTrue)
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sByte, tByte := []byte(s), []byte(t)
	m := make(map[byte]int)
	for i := 0; i < len(sByte); i++ {
		m[sByte[i]]++
	}
	for i := 0; i < len(tByte); i++ {
		m[tByte[i]]--
		if m[tByte[i]] < 0 {
			return false
		}
	}
	return true
}
