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

func main() {
	n := 27
	fmt.Println(isPowerOfThree2(n))
}
