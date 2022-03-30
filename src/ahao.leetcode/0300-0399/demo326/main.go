package main

import "fmt"

//326. 3的幂

//循环
func isPowerOfThree2(num int) bool {
	for num >= 3 {
		if num%3 == 0 {
			num = num / 3
		} else {
			return false
		}
	}
	return num == 1
}

func isPowerOfThree(num int) bool {
	return isPowerOfN(num, 3)
}

func isPowerOfN(num int, n int) bool {
	if num < 1 {
		return false
	}
	for num != 1 {
		if num%n == 0 {
			num = num / n
		}
	}
	return num == 1
}

func main() {
	n := 27
	fmt.Println(isPowerOfThree2(n))
}
