package main

import (
	"fmt"

)

func main(){
	var value1 float64
	value1=1
	value2:=1.00000000000000000000000001
	if value1==value2{
		fmt.Println("x等于y")
	}
}