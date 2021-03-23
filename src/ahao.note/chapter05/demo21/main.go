package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	//启动一个goroutine用于生成随机数，函数返回一个通道用于获取随机数
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

//生成器
func main() {
	ch := GenerateIntA()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
