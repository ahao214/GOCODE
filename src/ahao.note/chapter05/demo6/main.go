package main

import "fmt"

func main() {
	one := make([]int, 3, 5)
	two := make([]int, 3)
	three := []int{10, 20, 5: 30}
	fmt.Println(one, len(one), cap(one))
	fmt.Println(two, len(two), cap(two))
	fmt.Println(three, len(three), cap(three))
}
