package main

import "fmt"

/**
 * 寻找最后的山峰
 * @param a int整型一维数组
 * @return int整型
 */
func solve(a []int) int {
	res := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i+1] > a[i] {
			res = i + 1
		}
	}
	return res
}

func main() {
	a := []int{9, 8, 7, 6}
	fmt.Println(solve(a))
}
