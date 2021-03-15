package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"log"
	"strconv"
)

func main() {
	s := "9"
	n, err := strconv.ParseInt(s, 10, 64) //使用外部变量
	if err != nil {
		log.Fatalln(err)
	} else if n < 0 || n > 10 {
		log.Fatalln("invalid number")
	}
	fmt.Println(n)
}
