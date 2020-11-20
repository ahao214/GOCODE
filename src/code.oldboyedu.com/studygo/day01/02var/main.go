package main


import "fmt"

//批量声明变量
var (
	name string
	age int
	isOK bool
)

func main(){
	name = "Tom"
	age = 24
	isOK = true

	//Go语言中，非全局变量声明了，必须使用，不然编译器报错
	fmt.Printf("name is :%s", name)
	fmt.Println(age)
	fmt.Println(isOK)

	// 简短变量声明,只能在函数里面使用
	n:=24
	m:= 43
	fmt.Println(n)
	fmt.Println(m)
}