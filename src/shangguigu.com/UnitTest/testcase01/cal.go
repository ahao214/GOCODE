package testcase01

//一个被测试的函数
func add(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(m int, n int) int {
	return m - n
}
