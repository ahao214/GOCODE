package main

import "unicode"

//520. 检测大写字母
//我们定义，在以下情况时，单词的大写用法是正确的：
//全部字母都是大写，比如 "USA" 。
//单词中所有字母都不是大写，比如 "leetcode" 。
//如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
//给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。

func detectCapitalUse(word string) bool {
	if len(word) >= 2 && unicode.IsLower(rune(word[0])) && unicode.IsLower(rune(word[1])) {
		return false
	}
	for i := 2; i < len(word); i++ {
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[1])) {
			return false
		}
	}
	return true
}

func detectCapitalUse2(word string) bool {
	capCount := 0
	for i := 0; i < len(word); i++ {
		if unicode.IsUpper(rune(word[i])) {
			capCount++
		}
	}
	//全是小写字母
	if capCount == 0 {
		return true
	}
	//全是大写字母
	if capCount == len(word) {
		return true
	}
	//首字母大写其他小写字母
	if capCount == 1 && unicode.IsUpper(rune(word[0])) {
		return true
	}
	return false
}