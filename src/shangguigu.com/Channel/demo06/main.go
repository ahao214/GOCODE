package main

import (
	"fmt"
	"time"
)

func main() {
	//使用select可以解决从管道取数据的阻塞问题

	//1.定义一个管道
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道，5个string
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法，在遍历管道的时候，如果不关闭，会阻塞导致死锁(deadlock)

	//在实际开发中，我们不好确定是什么时候关闭管道
	//可以使用select来解决这个问题
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan中取到的数据是：%d\n", v)
			time.Sleep(time.Second)
		case v := <-strChan:
			fmt.Printf("从strChan中取到的数据是：%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("取不到数据，不玩了")
			time.Sleep(time.Second)
			return
		}
	}
}
