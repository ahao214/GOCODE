package main

import (
	"fmt"
)

//切片
func main() {
	//创建容量为5的切片
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}
