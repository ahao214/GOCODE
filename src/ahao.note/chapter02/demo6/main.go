package main

import (
	"fmt"
)

func f() int {
	a := 0
	defer func(i int) {
		fmt.Println("defer i=", i)
	}(a)
	a++
	return a
}

//defer
func main() {
	//先进后出
	defer func() {
		fmt.Println("first")
	}()
	defer func() {
		fmt.Println("second")
	}()
	fmt.Println("function body")

	a := 0
	fmt.Println(a)
	return

	defer func() {
		fmt.Println("third")
	}()
}
