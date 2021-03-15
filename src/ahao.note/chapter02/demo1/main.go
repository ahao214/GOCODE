package main

import (
	"fmt"
)

//iota
const (
	x = iota //0
	y        //1
	z        //2
)

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	fmt.Println(c)
}
func main() {
	test(red)
	test(100)
	//x := 2
	//test(x)
}
