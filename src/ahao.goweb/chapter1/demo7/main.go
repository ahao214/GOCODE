package main

import (
	"fmt"
)

//值传递
func main() {
	n := 6
	m := 8
	fmt.Printf("交换前n的值是：%d\n", n)
	fmt.Printf("交换前m的值是：%d\n", m)
	exchange(n, m)
	fmt.Printf("交换后n的值是：%d\n", n)
	fmt.Printf("交换后m的值是：%d\n", m)
}

func exchange(x, y int) int {
	var tmp int
	tmp = x
	x = y
	y = tmp
	return tmp
}
