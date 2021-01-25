package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

//runtime.GOMAXPROCS
//Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码
func main() {
	//两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务
	//runtime.GOMAXPROCS(1)

	//将逻辑核心数设为2，此时两个任务并行执行
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
