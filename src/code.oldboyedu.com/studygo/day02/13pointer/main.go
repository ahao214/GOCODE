package main

import "fmt"

func main(){
	//1.&取地址
	n:=10
	p:=&n
	fmt.Println(p)
	fmt.Printf("%T\n",p)

	//2.根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n",m)
}