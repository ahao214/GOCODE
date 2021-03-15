package main

import (
	"fmt"
	"log"
)

func test() {
	defer fmt.Println("test.1")
	defer fmt.Println("test.2")
	panic("i am dead")
}

//panic
func main() {
	defer func() {
		log.Println(recover())
	}()
	test()
}
