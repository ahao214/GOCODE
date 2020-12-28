package main

import(
	"fmt"
)

type Circle struct{
	radius float64
}

func(c Circle) area() float64{
	return 3.14*c.radius*c.radius
}

func main(){
	//创建一个结构体
	var c Circle
	c.radius=4.0
	res:=c.area()
	fmt.Println("面积是："+res)

}