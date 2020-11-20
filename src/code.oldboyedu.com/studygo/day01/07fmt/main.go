package main

import "fmt"


func main(){
	var i = 100
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", i)

	var name = "Tome"
	fmt.Printf("%s\n", name)
	fmt.Printf("%v\n", name)
}