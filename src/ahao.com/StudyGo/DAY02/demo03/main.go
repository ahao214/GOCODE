package main

import (
	"fmt"
)

func main() {

	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println("s:", s)
	fmt.Println("data:", data)

	one := []int{0, 1, 2, 3, 8: 100} //通过初始化表达式构造，可使用索引号
	fmt.Println(one, len(one), cap(one))

	two := make([]int, 6, 8) //使用make创建，指定len和cap
	fmt.Println(two, len(two), cap(two))

	three := make([]int, 6) //省略cap，相当于cap=len
	fmt.Println(three, len(three), cap(three))

}
