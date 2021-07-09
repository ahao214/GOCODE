package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//蛮力法
func power1(d float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return d
	}
	result := 1.0
	if n > 0 {
		for i := 1; i <= n; i++ {
			result *= d
		}
		return result
	} else {
		for i := 1; i <= Abs(n); i++ {
			result = result / d
		}
	}
	return result
}

//计算一个数的n次方
func main() {
	fmt.Println("蛮力法")
	fmt.Println(power1(2, 3))
	fmt.Println(power1(-2, 3))
	fmt.Println(power1(2, -3))

}
