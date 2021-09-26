package main

import (
	"fmt"
)

func print(n int) {
	if n > 0 {
		print(n - 1)
		fmt.Print(n, "")
	}
}

//如何不使用循环输出1到100
func main() {
	print(100)
	fmt.Println()
}
