package main

import (
	"fmt"
)

//如何优雅的从通道循环取值
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//开启goroutine将0-100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//开启goroutine从ch1中接收值，并将值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 //通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	//在主goroutine中从ch2中接收值打印
	for i := range ch2 { //通道关闭后悔退出for range 循环
		fmt.Println(i)
	}
}
