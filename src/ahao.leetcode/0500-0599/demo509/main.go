package main

//509. 斐波那契数
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	tmp := make([]int, n+1)
	tmp[0] = 0
	tmp[1] = 1
	for i := 2; i <= n; i++ {
		tmp[i] = tmp[i-1] + tmp[i-2]
	}
	return tmp[n]
}
