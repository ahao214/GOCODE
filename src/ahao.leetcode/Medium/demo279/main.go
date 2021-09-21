package main

import (
	"fmt"
	"math"
)

//279. 完全平方数
func numSquare(n int) int {
	if isPrefectSquare(n) {
		return 1
	}
	if checkAnser(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPrefectSquare(j) {
			return 2
		}
	}
	return 3
}

//判断是否为完全平方数
func isPrefectSquare(n int) bool {
	sq := int(math.Floor(math.Sqrt(float64(n))))
	return sq*sq == n
}

//判断是否能表示为4^k*(8m+7)
func checkAnser(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

func main() {
	n := 8
	fmt.Println(numSquare(n))
}
