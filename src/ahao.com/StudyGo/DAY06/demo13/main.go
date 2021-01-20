package main

import (
	"fmt"
)

//Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

//dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

//car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s的时速是70迈\n", c.brand)
}

//多个类型实现同一接口
func main() {
	var m Mover
	var d = dog{name: "旺财"}
	var c = car{brand: "法拉利"}
	m = d
	m.move()
	m = c
	m.move()
}
