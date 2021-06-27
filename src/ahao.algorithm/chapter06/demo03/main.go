package main

import (
	"fmt"
)

//减法
//方法功能：计算两个自然数的除法
func divide(m, n int) (res, remain int) {
	res = 0
	remain = m
	//被除数减除数，直到相减结果小于除数为止
	for m > n {
		m = m - n
		res++
	}
	remain = m
	return
}

//移位法
func divide2(m, n int) (res, remain int) {
	res = 0
	for m >= n {
		multi := 1
		//multi*n>m/2时结束循环
		for multi*n <= (m >> 1) {
			multi <<= 1
		}
		res += multi
		//相减的结果进入下次循环
		m -= multi * n
	}
	remain = m
	return
}

//不使用加减乘除运算实现加法
func add(m, n int) int {
	//保存不进位相加结果
	sum := 0
	//保存进位值
	carry := -1
	for carry != 0 { //判断进位值是否为0
		sum = n ^ m          //异或代替不进位相加
		carry = (n & m) << 1 //与操作代替计算进位值
		n = sum
		m = carry
	}
	return sum
}

//不使用除法操作符实现两个正整数的除法
func main() {
	m := 14
	n := 4
	fmt.Println("减法：")
	fmt.Println(m, "除以", n)
	res, remain := divide(m, n)
	fmt.Println("商为：", res, "余数为：", remain)

	fmt.Println("移位法：")
	fmt.Println("减法：")
	fmt.Println(m, "除以", n)
	res, remain = divide2(m, n)
	fmt.Println("商为：", res, "余数为：", remain)

	fmt.Println("不使用加减乘除运算实现加法")
	fmt.Println(add(2, 4))

}
