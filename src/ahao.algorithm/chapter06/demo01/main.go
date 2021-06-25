package main

import (
	"fmt"
)

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

//二分查找法
func isPower2(n int) bool {
	if n <= 0 {
		fmt.Println(n, "不是自然数")
		return false
	}
	low := 0
	high := n
	for low < high {
		mid := (low + high) / 2
		power := mid * mid
		if power > n { //大于则在小的范围找
			high = mid - 1
		} else if power < n { //小于则在打的范围找
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}

//减法运算法
func isPower3(n int) bool {
	if n <= 0 {
		fmt.Println(n, "不是自然数")
		return false
	}
	minus := 1
	for n > 0 {
		n = n - minus
		//n是某个数的平方
		if n == 0 {
			return true
		} else if n < 0 {
			return false
		} else {
			minus += 2
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

	fmt.Println("二分查找法：")
	if isPower2(n1) {
		fmt.Println(n1, "是某个自然数的平方")
	} else {
		fmt.Println(n1, "不是某个自然数的平方")
	}

	if isPower2(n2) {
		fmt.Println(n2, "是某个自然数的平方")
	} else {
		fmt.Println(n2, "不是某个自然数的平方")
	}

	fmt.Println("减法运算法：")
	if isPower3(n1) {
		fmt.Println(n1, "是某个自然数的平方")
	} else {
		fmt.Println(n1, "不是某个自然数的平方")
	}

	if isPower3(n2) {
		fmt.Println(n2, "是某个自然数的平方")
	} else {
		fmt.Println(n2, "不是某个自然数的平方")
	}
}
