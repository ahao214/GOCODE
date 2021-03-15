package main

import (
	"fmt"
)

func test() []func() {
	var s []func()
	for i := 0; i < 2; i++ {
		s = append(s, func() { //将多个匿名函数添加到列表
			fmt.Println(&i, i)
		})
	}
	return s
}

//闭包
func main() {
	for _, f := range test() { //迭代执行所有匿名函数
		f()
	}
}
