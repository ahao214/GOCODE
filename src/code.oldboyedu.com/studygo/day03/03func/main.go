package main

import "fmt"


func fone(){
	fmt.Println("hello world")
}

func f2(name string){
	fmt.Println("hello,",name)

}

func f3(x int,y int) int{
	sum:= x + y
	return sum
}

func main(){
	fone()
	f2("深圳")
	ret:=f3(100,111)
	fmt.Println(ret)

}