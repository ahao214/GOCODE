package main

import (
	"fmt"
	"time"
)

func longWait() {
	fmt.Println("开始longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("结束longWait()")
}

func shortWait() {
	fmt.Println("开始shortWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("结束shortWait()")
}

func main() {
	fmt.Println("这里是main()开始的地方")
	go longWait()
	go shortWait()
	fmt.Println("挂起main()")
	//挂起工作时间以纳秒(ns)为单位
	time.Sleep(10 * 1e9)
	fmt.Println("这里是main()结束的地方")
}
