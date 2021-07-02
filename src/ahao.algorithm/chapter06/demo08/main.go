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

//数字规律法
func serach2(n int) int {
	a := []int{0, 2, 3, 4, 5, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28, 30}
	return (n/22)*30 + a[n%22]
}

//求有序数列的第1500个数的值
func main() {
	fmt.Println("蛮力方法")
	fmt.Println(serach1(1500))

	fmt.Println("数字规律法")
	fmt.Println(serach2(1500))
}
