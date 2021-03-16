package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	one := append(s, 10)
	//追加多个数据
	two := append(s, 20, 30)

	fmt.Println(s, len(s), cap(s))
	fmt.Println(one, len(one), cap(one))
	fmt.Println(two, len(two), cap(two))
}
