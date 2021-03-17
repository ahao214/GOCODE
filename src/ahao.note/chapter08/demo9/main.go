package main

import (
	"fmt"
)

//通道
func main() {
	done := make(chan struct {
	})
	//数据传输通道
	c := make(chan string)
	go func() {
		s := <-c //接收消息
		fmt.Println(s)
		close(done)
	}()
	c <- "nil" //发送消息
	<-done

}
