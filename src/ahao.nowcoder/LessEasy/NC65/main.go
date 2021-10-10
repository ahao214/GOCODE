package main

import "fmt"

/**
 * 斐波那契数列
 * @param n int整型
 * @return int整型
 */
func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(Fibonacci(3))
}
