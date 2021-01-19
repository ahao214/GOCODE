package main

import (
	"fmt"
)

//定义相互交换值的函数
//引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
func swap(x, y *int) {
	var temp int
	temp = *x //保存x的值
	*x = *y   //将y值赋给x
	*y = temp //将temp值赋给y
}

//值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
func swapone(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}

//不定参数传值
func testone(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

//使用slice对象做变参时，必须展开
func testtwo(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func main() {

	var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	fmt.Println(a, b)
	var x, y int = 3, 4
	swapone(x, y)
	fmt.Println(x, y)

	fmt.Println(testone("sum:%d", 1, 2, 3))

	s := []int{4, 5, 6}
	res := testtwo("sum:%d", s...) // slice... 展开slice
	fmt.Println(res)

}
