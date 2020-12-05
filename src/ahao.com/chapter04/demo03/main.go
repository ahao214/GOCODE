package main

import(
	"fmt"
)

//select语句
//select是随机选择一个case来判断，直到匹配其中的一个case


func main(){
	a:=make(chan int,1024)
	b:=make(chan int,1024)
	for i:=0;i<10;i++{
		fmt.Printf("第%d次，",i)
		a <- 1
		b <- 1
		select{
		case <- a:
			fmt.Println("from a")
		case <- b:
			fmt.Println("from b")
		}
	}

}