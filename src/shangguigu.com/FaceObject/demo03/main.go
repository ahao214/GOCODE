package main

import(
	"fmt"
)
//结构体
type Point struct{
	x int
	y int
}

//结构体
type Rect struct{
	leftUp,rightDown Point
}



func main(){

	r1:=Rect{Point{1,2},Point{3,4}}
	//r1有四个int，在内存中是连续分布
	//打印地址
	fmt.Printf("r1.leftUp.x的地址是%p\nr1.leftUp.y的地址是%p\nr1.rightDown.x的地址是%p\nr1.rightDown.y的地址是%p\n",&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)




}