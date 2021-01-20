package main

import (
	"fmt"
)

//延迟调用defer
func main() {
	//defer是先进后出
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i)
	}

	//defer碰上闭包
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}

}
