package main

import(
	"fmt"
	"time"
)


//goroutine
func hello(i int){
	fmt.Println("hello",i)
}

func main(){
	for i:=0;i<100;i++{
		// go hello(i)	//开启一个单独的goroutine去执行hello函数(任务)
		go func(i int){
			fmt.Println(i)
		}(i)
	}
	
	fmt.Println("main")
	time.Sleep(time.Second)
}