package main

import (
	"errors"
	"fmt"
	"log"
)

var errDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errDivByZero
	}
	return x / y, nil
}

//错误处理
func main() {
	z, err := div(5, 0)
	if err == errDivByZero {
		log.Fatalln(err)
	}
	fmt.Println(z)
}
