package main

//1422. 分割字符串的最大得分
func maxScore(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	left := 0
	if s[0] == '0' {
		left++
	}
	right := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			right++
		}
	}
	res = left + right
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			left++
		}
		if s[i] == '1' {
			right--
		}
		tmp := left + right
		if tmp > res {
			res = tmp
		}
	}
	return res
}
