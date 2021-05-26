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

//二分法
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	left, right, res := 1, x, 0
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid < x/mid {
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
		}
	}
	return res
}

func main() {
	x := 8
	fmt.Println("牛顿迭代法")
	fmt.Println(mySqrt(x))
	fmt.Println("二分法")
	fmt.Println(mySqrt2(x))
}
