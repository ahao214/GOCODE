package main

import "fmt"

//X的平方根

//牛顿迭代法
func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}
