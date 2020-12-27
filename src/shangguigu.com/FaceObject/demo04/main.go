package main

import(
	"fmt"
)

type A struct{
	Num int
}

type B struct{
	Num int
}

func main(){
	var a A
	var b B
	//两个结构体的字段完全相同(字段的名字、个数、类型要完全一样)，就可以进行转换
	a=A(b)
	fmt.Println(a,b)



}