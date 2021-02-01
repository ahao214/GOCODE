package main

import (
	"fmt"
)

func Add(a int) int {
	a++
	return a
}

func AddP(a *int) int {
	*a++
	return *a
}

func main() {
	x := 3
	fmt.Println("x=", x, "&x=", &x)
	y := Add(x)                   //执行Add函数，但实际上修改的是x的副本
	fmt.Println("x=", x, "y=", y) //输出的x还是原来的值，改变的是副本y的值
	z := AddP(&x)                 //调用函数AddP，实际上修改的是x的值
	fmt.Println("x=", x, "z=", z) //输出的x已经被修改
	fmt.Println("&x =", &x)
}
