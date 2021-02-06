package main

import (
	"fmt"
)

type S struct {
	Name string
}

func (s S) M1() {
	s.Name = "value"
}
func (s S) M2() {
	s.Name = "pointer"
}
func main() {
	var sone = S{"new"}
	var stwo = &sone
	sone.M2()
	fmt.Printf("%v,%v\n", sone, stwo)

	sone = S{"new"}
	stwo = &sone
	stwo.M1()
	fmt.Printf("%v,%v\n", sone, stwo)
}
