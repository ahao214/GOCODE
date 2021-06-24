package main

import "fmt"

//直接计算法
func isPower(n int) bool {
	if n <= 0 {
		fmt.Println(n, "不是自然数")
		return false
	}
	for i := 1; i < n; i++ {
		m := i * i
		if m == n {
			return true
		} else if m > n {
			return false
		}
	}
	return false
}

//判断一个自然数是否是某个数的平方
func main() {
	n1 := 15
	n2 := 16
	fmt.Println("直接计算法：")
	if isPower(n1) {
		fmt.Println(n1, "是某个自然数的平方")
	} else {
		fmt.Println(n1, "不是某个自然数的平方")
	}

	if isPower(n2) {
		fmt.Println(n2, "是某个自然数的平方")
	} else {
		fmt.Println(n2, "不是某个自然数的平方")
	}

}
