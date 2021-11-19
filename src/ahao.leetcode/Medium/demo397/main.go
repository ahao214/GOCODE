package main

//397. 整数替换

/*
枚举所有的情况
*/
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
