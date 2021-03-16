package main

import "fmt"

type N int

func (n N) test() {
	fmt.Println("test:n:%p,%v\n", &n, n)
}

func main() {
	var n N = 100
	p := &n
	n++
	f1 := n.test
	n++
	f2 := p.test

	n++
	fmt.Printf("main.n:%p,%v\n", p, n)
	f1()
	f2()
}
