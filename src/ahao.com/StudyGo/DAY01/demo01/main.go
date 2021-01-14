package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 100
	var b int = 200
	var t int
	t = a
	a = b
	b = t
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("t=", t)

	//按照默认宽度和精度输出整型
	fmt.Printf("%f\n", math.Pi)
	//按默认宽度，2位精度输出
	fmt.Printf("%.2f\n", math.Pi)

	q := make([]int, 3)
	q[0] = 1
	q[1] = 2
	q[2] = 3
	str := "hello world"
	fmt.Println(str[6:])

}
