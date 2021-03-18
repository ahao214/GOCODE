package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("func body")
	os.Exit(1)
}
