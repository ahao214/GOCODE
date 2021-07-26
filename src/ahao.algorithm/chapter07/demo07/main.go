package main

import (
	"fmt"
)

func factorIsOdd(a int) int {
	var total int
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			total++
		}
	}
	if total%2 == 1 {
		return 1
	} else {
		return 0
	}
}

func totalCount(num []int, n int) int {
	var count int
	for i := 0; i < n; i++ {
		if factorIsOdd(num[i]) == 1 {
			fmt.Println("亮着的灯的编号是：", num[i])
			count++
		}
	}
	return count
}

//如何判断还有几盏灯泡还亮着
func main() {
	num := make([]int, 100)
	for i := 0; i < 100; i++ {
		num[i] = i + 1
	}
	count := totalCount(num, 100)
	fmt.Println("最后总共有,", count, "盏灯亮着")
}
