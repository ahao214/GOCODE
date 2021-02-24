package main

import "fmt"

//第231题：2的幂
func main() {
	result := isPowerOfTwo(9)
	fmt.Println(result)
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
