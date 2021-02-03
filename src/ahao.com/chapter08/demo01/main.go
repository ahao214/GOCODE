package main

import (
	"fmt"
)

type myStruct struct {
	i1  int
	f1  float64
	str string
}

func main() {
	ms := new(myStruct)
	ms.i1 = 12
	ms.f1 = 15.1
	ms.str = "hello Go"
	fmt.Printf("int:%d\n", ms.i1)
	fmt.Printf("float:%f\n", ms.f1)
	fmt.Printf("string:%s\n", ms.str)
	fmt.Println(ms)
}
