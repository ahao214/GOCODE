package main

/*
1422. 分割字符串的最大得分
*/


func maxScore(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		ans = max(ans, strings.Count(s[:i], "0")+strings.Count(s[i:], "1"))
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}


//两次遍历
func maxScore2(s string) int {
	score := int('1'-s[0]) + strings.Count(s[1:], "1")
	ans := score
	for _, c := range s[1 : len(s)-1] {
		if c == '0' {
			score++
		} else {
			score--
		}
		ans = max(ans, score)
	}
	return ans
}

