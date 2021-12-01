package main

//1446. 连续字符
func maxPower(s string) int {
	result := 1
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			result = max(result, count)
		} else {
			count = 1
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
