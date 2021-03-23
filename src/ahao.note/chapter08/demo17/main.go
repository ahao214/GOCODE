package main

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

//向方法传递指针引用
func main() {
	t := Triangle{
		base:   3,
		height: 1,
	}
	fmt.Println(t.area())
}
