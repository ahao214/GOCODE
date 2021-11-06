package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 找缺失数字
 * @param a int整型一维数组 给定的数字串
 * @return int整型
 */
func solve(a []int) int {
	result := 0
	for i := 0; i < len(a); i++ {
		if a[i] != i {
			result = i
			break
		}
		result = i + 1
	}
	return result
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 7}
	fmt.Println(solve(a))
}
