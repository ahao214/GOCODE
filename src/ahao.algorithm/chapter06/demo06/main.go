package main

import (
	"fmt"
)

func zeroCount(n int) int {
	count := 0
	for n > 0 {
		n = n / 5
		count += n
	}
	return count
}

//判断1024!末尾有多少个0
func main() {
	fmt.Println("1024!末尾0的个数为：", zeroCount(1024))
}
