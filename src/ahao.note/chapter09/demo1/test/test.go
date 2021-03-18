package main

import (
	"ahao.note/chapter09/demo1/lib"
	"fmt"
	"unsafe"
)

func main() {
	lib.Hello()

	d := lib.NewData()
	d.Y = 100
	p := (*struct {
		x int
	})(unsafe.Pointer(d))
	p.x = 200
	fmt.Printf("%+v\n", *d)

}
