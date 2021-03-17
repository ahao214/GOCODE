package main

import (
	"fmt"
)

//缓冲区大小和当前已缓冲数量
func main() {
	a, b := make(chan int), make(chan int, 3)
	b <- 1
	b <- 2
	fmt.Println("a:", len(a), cap(a))
	fmt.Println("b:", len(b), cap(b))

	fmt.Println("--------")

	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for x := range c { //循环获取消息，直到通道被关闭
			fmt.Println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
