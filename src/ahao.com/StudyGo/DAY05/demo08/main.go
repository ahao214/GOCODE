package main

import "fmt"

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func test() func() {
	x := 100
	fmt.Printf("x (%p)=%d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

//返回2个函数类型的返回值
func testone(base int) (func(int) int, func(int) int) {
	//定义2个函数，并返回
	//相加
	add := func(i int) int {
		base += i
		return base
	}
	//相减
	sub := func(i int) int {
		base -= i
		return base
	}
	//返回
	return add, sub
}

func main() {
	c := a()
	c()
	c()
	c()

	f := test()
	f()

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	//此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	add, sub := testone(10)
	//base一直没有消
	fmt.Println(add(1), sub(2))
	// 此时base是9
	fmt.Println(add(3), sub(4))

}
