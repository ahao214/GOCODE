package main

//1816. 截断句子
func truncateSentence(s string, k int) string {
	n, end, count := len(s), 0, 0
	for i := 1; i <= n; i++ {
		if i == n || s[i] == ' ' {
			count++
			if count == k {
				end = i
				break
			}
		}
	}
	return s[:end]
}
