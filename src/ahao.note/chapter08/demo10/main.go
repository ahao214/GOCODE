package main

import (
	"fmt"
	"unsafe"
)

//缓冲
func main() {
	//创建带3个缓冲槽的异步通道
	a := make(chan int, 3)
	a <- 1
	a <- 2
	fmt.Println(<-a)
	fmt.Println(<-a)

	var x, y chan int = make(chan int, 3), make(chan int)
	var c chan bool
	fmt.Println(x == y)
	fmt.Println(c == nil)
	fmt.Printf("%p,%d\n", a, unsafe.Sizeof(x))
}
