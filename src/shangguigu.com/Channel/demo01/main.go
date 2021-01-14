package main

import (
	"fmt"
)

func main() {
	//演示一下管道的使用
	//1.创建一个可以存在3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	//2.看看intChan是什么
	fmt.Printf("intChan的值是：%v，intChan本身的地址是:%p\n", intChan, &intChan)

	//3.向管道写入数据
	intChan <- 10
	num := 20
	intChan <- num
	//当给管道写入数据的时候，不能超过管道的容量
	//4.查看管道的长度和容量
	fmt.Printf("chan len=%v, cap=%v\n", len(intChan), cap(intChan))

	//从管道中读取数据
	var one int
	one = <-intChan
	fmt.Println(one)
	fmt.Printf("chan len=%v, cap=%v\n", len(intChan), cap(intChan))

}
