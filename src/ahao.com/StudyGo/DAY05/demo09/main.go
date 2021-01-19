package main

import (
	"fmt"
)

//数字阶乘
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

//斐波那契数列(Fibonacci)
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

//递归函数
func main() {

	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("%d\n", fibonaci(j))
	}

}
