package main

import (
	"fmt"
)

//蛮力方法
func serach1(n int) int {
	i, count := 0, 0
	for true {
		i++
		if i%2 == 0 || i%3 == 0 || i%5 == 0 {
			count++
		}
		if count == n {
			break
		}
	}
	return i
}

//求有序数列的第1500个数的值
func main() {
	fmt.Println("蛮力方法")
	fmt.Println(serach1(1500))
}
