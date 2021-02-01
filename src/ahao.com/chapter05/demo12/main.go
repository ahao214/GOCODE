package main

import (
	"fmt"
)

func multiply(a, b int, res *int) {
	*res = a * b
}

func main() {
	n := 0
	res := &n
	multiply(2, 4, res)
	fmt.Println("result:", *res)
}
