package main

import (
	"fmt"
)

//数组
func main() {
	var a [4]int
	b := [4]int{2, 5}
	c := [4]int{5, 3: 10}
	d := [...]int{1, 2, 3}
	e := [...]int{10, 3: 100}
	fmt.Println(a, b, c, d, e)

	type user struct {
		name string
		age  byte
	}

	g := [...]user{
		{"Tom", 20},
		{"Mary", 18},
	}
	fmt.Printf("%#v\n", g)

}
