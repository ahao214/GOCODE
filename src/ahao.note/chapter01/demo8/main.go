package main

import (
	"fmt"
)

//消费者
func consumer(data chan int, done chan bool) {
	for x := range data { //接收数据，直到通道被关闭
		fmt.Println("recv:", x)
	}
	//通知main，消费结束
	done <- true
}

//生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i //发送数据
	}
	close(data) //生产结束，关闭通道
}

//通道(channel)与goroutine搭配，实现用通信代替内存共享的CSP模型
func main() {
	done := make(chan bool) //用于接收消费结束信号
	data := make(chan int)  //数据管道

	go consumer(data, done) //启动消费者
	go producer(data)       //启动生产者

	<-data //阻塞，直到消费者发回结束信号

}
