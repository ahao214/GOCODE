package main

import "fmt"

func adder() func(int)int{
	var x int
	return func(y int) int{
		x+=y
		return x
	}
}


func adder2(x int) func(int) int{
	return func(y int)int{
		x+=y
		return x
	}
}


func main(){
	ret:=adder()
	ret2:=ret(200)
	fmt.Println(ret2)

	ret3:=adder2(300)
	ret4:=ret3(400)
	fmt.Println(ret4)
}