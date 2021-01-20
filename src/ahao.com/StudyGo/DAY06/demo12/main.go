package main

import (
	"fmt"
)

//Sayer接口
type Sayer interface {
	say()
}

//Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

//实现sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

//实现mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

//一个类型实现多个接口
func main() {
	var x Sayer
	var y Mover
	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()

}
