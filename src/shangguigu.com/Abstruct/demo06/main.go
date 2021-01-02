package main

import(
	"fmt"
)

type A struct{
	Name string
	age int
}

type B struct{
	Name string
	Score float64
}

type C struct{
	A
	B
}

type D struct{
	a A
}

func main(){

	var c C
	c.A.Name="jack"

	fmt.Println(c.A.Name)

	var d D
	d.a.Name="sky"
	fmt.Println(d.a.Name)
}