package main

import (
	"fmt"
)

//全局变量m
var m = 100

//匿名变量
func foo() (int, string) {
	return 10, "kevin"
}

func main() {
	n := 10
	m := 20
	fmt.Println(n, m)

	x, _ := foo()
	_, z := foo()
	fmt.Println("x=", x)
	fmt.Println("z=", z)

}
