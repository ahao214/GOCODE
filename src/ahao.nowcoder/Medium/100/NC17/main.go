package main

//NC17 最长回文子串
func getLongestPalindrome(A string, n int) int {
	if A == "" {
		return 0
	}
	start, end := 0, 0
	for i := 0; i < len(A); i++ {
		left1, right1 := expandAroundCenter(A, i, i)
		left2, right2 := expandAroundCenter(A, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return end - start + 1
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
