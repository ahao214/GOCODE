package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func GenerateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func GenerateInt() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			//使用select的扇入技术(Fan in)增加生产的随机源
			select {
			case ch <- <-GenerateIntA():
			case ch <- <-GenerateIntB():
			}
		}
	}()
	return ch
}

//多个goroutine增强型生成器
func main() {
	ch := GenerateInt()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
