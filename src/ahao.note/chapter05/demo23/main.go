package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			//通过select监听一个信号chan来确定是否停止生成
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//不再需要生成器，通过close chan发送一个通知给生成器
	close(done)
	for v := range ch {
		fmt.Println(v)
	}
}
