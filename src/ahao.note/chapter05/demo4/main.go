package main

import (
	"fmt"
)

//定义多维数组时，仅第一维度允许使用...
func main() {
	a := [2][2]int{
		{1, 2},
		{3, 4},
	}
	b := [...][2]int{
		{10, 20},
		{30, 40},
	}
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	fmt.Println(len(b[1]), cap(b[1]))
}
