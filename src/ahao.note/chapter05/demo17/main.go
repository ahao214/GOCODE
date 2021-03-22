package main

import (
	"fmt"
)

//chan
func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		//写通道
		c <- struct{}{}
	}(c, ci)

	for v := range ci {
		fmt.Println(v)
	}
}
