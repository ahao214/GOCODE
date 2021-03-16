package main

import (
	"fmt"
)

func main() {
	x := [][]int{
		{1, 2},
		{10, 20, 30},
		{100},
	}
	fmt.Println(x)

	x[2] = append(x[2], 200, 300)
	fmt.Println(x[2])
}
