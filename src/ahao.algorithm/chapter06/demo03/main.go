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

//不使用除法操作符实现两个正整数的除法
func main() {
	m := 14
	n := 4
	fmt.Println("减法：")
	fmt.Println(m, "除以", n)
	res, remain := divide(m, n)
	fmt.Println("商为：", res, "余数为：", remain)

}
