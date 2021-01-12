package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写一个函数，每隔1秒输出"hello world"
func test(){
	for i:=1;i<=10;i++{
		fmt.Println("hello world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main(){
	go test()	//开启了一个协程
	for i:=1;i<=10;i++{
		fmt.Println("main hello golang"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
