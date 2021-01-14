package main

import (
	"fmt"
)

//切片初始化
var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var zero []int = arr[2:8]
var one []int = arr[0:6]
var two []int = arr[5:10]
var three []int = arr[0:len(arr)]
var four = arr[:len(arr)-1]

func main() {

	fmt.Printf("全局变量：arr%v\n", arr)
	fmt.Printf("全局变量：zero%v\n", zero)
	fmt.Printf("全局变量：one%v\n", one)
	fmt.Printf("全局变量：two%v\n", two)
	fmt.Printf("全局变量：three%v\n", three)
	fmt.Printf("全局变量：four%v\n", four)
	fmt.Printf("-----------------\n")

	arrtwo := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	five := arrtwo[2:8]
	six := arrtwo[0:6]
	seven := arrtwo[5:10]
	fmt.Printf("局部变量：five%v\n", five)
	fmt.Printf("局部变量：six%v\n", six)
	fmt.Printf("局部变量：seven%v\n", seven)

}
