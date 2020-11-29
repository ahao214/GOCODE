package main

import(
	"fmt"
)

var x int32
var y int64


func main(){
	x,y=2,4
	if x==y{ //编译不通过
		fmt.Println("x等于y")
	}
	if x==2 || y==2{
		fmt.Println("x等于y")
	}

}