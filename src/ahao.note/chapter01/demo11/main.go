package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	*Person
	Number int
}

func main() {
	a := Person{"Tom", 21}
	fmt.Println(a)

	p := &Person{
		Name: "Kevin",
		Age:  29,
	}
	s := Student{
		Person: p,
		Number: 110,
	}
	fmt.Println(s)

}
