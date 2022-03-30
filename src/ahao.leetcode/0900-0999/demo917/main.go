package main

import "unicode"

/*
917. 仅仅反转字母
给你一个字符串 s ，根据下述规则反转字符串：

所有非英文字母保留在原有位置。
所有英文字母（小写或大写）位置反转。
返回反转后的 s 。
*/
func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	left, right := 0, len(s)-1
	for {
		for left < right && !unicode.IsLetter(rune(s[left])) {
			left++
		}
		for right > left && !unicode.IsLetter(rune(s[right])) {
			right--
		}
		if left >= right {
			break
		}
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return string(ans)
}
