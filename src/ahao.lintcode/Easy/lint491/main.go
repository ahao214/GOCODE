package main

import "strconv"

/**
 * 491.回文数
 * @param num: a positive number
 * @return: true if it's a palindrome or false
 */
func isPalindrome(num int) bool {
	if num%10 == 0 && num != 0 {
		return false
	}
	n := strconv.Itoa(num)
	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-i-1] {
			return false
		}
	}
	return true
}
