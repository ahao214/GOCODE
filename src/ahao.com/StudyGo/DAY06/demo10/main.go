package main

import (
	"fmt"
)

type Mover interface {
	move()
}

type dog struct {
}

//指针接受者实现接口
func (d *dog) move() {
	fmt.Println("狗会跑动")
}

func main() {
	var x Mover
	//var wangcai=dog{}
	//x=wangcai	//x不可以接收dog类型

	var fugui = &dog{}
	x = fugui
	x.move()

}
