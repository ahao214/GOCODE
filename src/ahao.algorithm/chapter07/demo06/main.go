package main

import (
	"fmt"
)

//蛮力法
func combinationCount(n int) int {
	count := 0
	num1 := n
	num2 := n / 2
	num3 := n / 5
	for x := 0; x <= num1; x++ {
		for y := 0; y <= num2; y++ {
			for z := 0; z <= num3; z++ {
				if x+2*y+5*z == n {
					count++
				}
			}
		}
	}
	return count
}

//数字规律法
func combinationCount2(n int) int {
	count := 0
	for m := 0; m <= n; m += 5 {
		count += (m + 2) / 2
	}
	return count
}

//如何组合1，2,5这三个数使其和为100
func main() {
	fmt.Println("蛮力法：")
	fmt.Println(combinationCount(100))

	fmt.Println("数字规律法：")
	fmt.Println(combinationCount2(100))
}
