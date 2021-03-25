package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	_, err := os.Getwd()
	if err != nil {
		log.Fatal("os.Getwd failed:%v", err)
	}
}

func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
}
