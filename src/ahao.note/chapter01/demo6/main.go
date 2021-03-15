package main

import (
	"fmt"
)

type X int

func (x *X) inc() {
	*x++
}

//方法
func main() {
	var x X
	x.inc()
	fmt.Println(x)
}
