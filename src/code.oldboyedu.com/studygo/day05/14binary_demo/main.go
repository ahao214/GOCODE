package main

import "fmt"

const (
	eat int = 4
	sleep int=2
	da int=1

)
//二进制实用途
//11
//最左边的1表示吃饭
//中间的1表示睡觉
//最右边的1表示打豆豆


func f(arg int){
	fmt.Printf("%b\n",arg)
}


func main(){
	f(eat | da)
}