package main

import(
	"fmt"
)

var name="tom"

func test01(){
	fmt.Println(name)
}

func test02(){
	name:="jack"
	fmt.Println(name)
}

func main(){
	fmt.Println(name) 	//tom
	test01()			//tom
	test02()			//jack
	test01()			//tom
}