package main

import (
	"fmt"
)

//获取n的平方根，e为精度要求
func squareRoot(n, e float64) float64 {
	newOne := n
	lastOne := 1.0           //第一个近似值为1
	for newOne-lastOne > e { //直到满足精度要求为止
		newOne = (newOne + lastOne) //求下一个近似值
		lastOne = n / newOne
	}
	return newOne
}

//在不能使用库函数的条件下计算n的算术平方根
func main() {
	n := 50.0
	e := 0.000001
	fmt.Printf("%f的平方根为%f", n, squareRoot(n, e))
	fmt.Println()
	n = 4.0
	fmt.Printf("%f的平方根为%f", n, squareRoot(n, e))
	fmt.Println()
}
