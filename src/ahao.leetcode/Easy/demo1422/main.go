package main

//1422. 分割字符串的最大得分
/*
给你一个由若干 0 和 1 组成的字符串 s ，请你计算并返回将该字符串分割成两个 非空 子字符串（即 左 子字符串和 右 子字符串）所能获得的最大得分。

「分割字符串的得分」为左子字符串中0的数量加上右子字符串中1的数量。
输入：s = "00111"
输出：5
解释：当 左子字符串 = "00" 且 右子字符串 = "111" 时，我们得到最大得分 = 2 + 3 = 5
*/
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
