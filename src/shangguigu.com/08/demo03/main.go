package main

import(
	"fmt"
)

//累加器
func AddUpper() func(int)int{
	var n int = 10
	var str ="hello"
	return func(x int)int{
		n = n + x
		str += "a"
		fmt.Println("str=",str)
		return n
	}
}

//闭包
func main(){
	f:=AddUpper()
	fmt.Println(f(1))	//输出11
	fmt.Println(f(2))	//输出13
	fmt.Println(f(3))	//输出16

}