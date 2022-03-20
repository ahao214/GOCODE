package main

/*
125. 验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
*/
func isPalindrome(s string) bool {
	//双指针
	if len(s) == 0 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isLetterOrNumber(s[left]) {
			left++
		}
		for left < right && !isLetterOrNumber(s[right]) {
			right--
		}
		if left <= len(s)-1 && right >= 0 {
			if toUpper(s[left]) != toUpper(s[right]) {
				return false
			}
		}
		left++
		right--
	}
	return true
}

func isLetterOrNumber(char byte) bool {
	return isUpper(char) || isLower(char) || isNumber(char)
}

func isUpper(char byte) bool {
	return 65 <= char && char <= 90
}

func isLower(char byte) bool {
	return 97 <= char && char <= 122
}

func isNumber(char byte) bool {
	return 48 <= char && char <= 57
}

func toUpper(char byte) byte {
	if isLower(char) {
		return byte(char - 32)
	}
	return char
}
