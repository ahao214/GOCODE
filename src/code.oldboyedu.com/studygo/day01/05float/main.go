package main

import "fmt"

func main(){
	//浮点数
	f1 := 1.234567
	fmt.Printf("%T\n", f1)	//Go语言中，默认的浮点数类型是64位
	f2 :=float32(1.234566) // 显示声明float32位
	fmt.Printf("%T\n",f2)
}