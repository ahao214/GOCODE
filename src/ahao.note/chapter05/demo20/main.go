package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

//GenerateIntA是一个随机数生成器
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		//收到通知后关闭通道ch
		close(ch)
	}()
	return ch
}

//通知退出机制
func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//发送通知，告诉生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//此时生产者已经退出
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())

}
