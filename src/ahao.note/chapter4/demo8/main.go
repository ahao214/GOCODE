package main

import (
	"fmt"
	"log"
)

//recover
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()
	panic("i am dead")
	fmt.Println("exit...")
}
