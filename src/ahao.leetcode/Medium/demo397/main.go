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

/*
记忆化搜索
*/
var dic map[int]int

func intReplacement(n int) (res int) {
	if n == 1 {
		return 0
	}
	if res, ok := dic[n]; ok {
		return res
	}
	defer func() { dic[n] = res }()
	if n%2 == 0 {
		return 1 + intReplacement(n/2)
	}
	return 2 + min(intReplacement(n/2), intReplacement(n/2+1))
}

func integerReplacement2(n int) int {
	dic = map[int]int{}
	return intReplacement(n)
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
