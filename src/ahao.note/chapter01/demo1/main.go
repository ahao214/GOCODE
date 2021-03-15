package main

import (
	"fmt"
)

//三种流程控制语句
func main() {
	x := 100
	if x > 0 {
		fmt.Println("x")
	} else if x < 0 {
		fmt.Println("-x")
	} else {
		fmt.Println(0)
	}

	y := 10
	switch {
	case y > 0:
		fmt.Println("x")
	case y < 0:
		fmt.Println("-x")
	default:
		fmt.Println("0")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	z := 0
	for z < 5 {
		fmt.Println(z)
		z++
	}

	m := []int{1, 2, 3}
	for i, n := range m {
		fmt.Println(i, ":", n)
	}

}
