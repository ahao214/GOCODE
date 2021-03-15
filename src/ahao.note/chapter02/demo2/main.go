package main

import (
	"fmt"
)

func main() {
	type data [2]int
	var d data = [2]int{1, 2}
	fmt.Println(d)
	a := make(chan int, 2)
	var b chan<- int = a
	b <- 2

}
