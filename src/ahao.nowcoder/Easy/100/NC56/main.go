package main

import "strconv"

/**
 * 回文数字
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	if (x%10 == 0 && x != 0) || x < 0 {
		return false
	}

	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
