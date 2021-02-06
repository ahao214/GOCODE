package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

func main() {
	p := new(Person)
	p.SetFirstName("Tom")
	fmt.Println(p.FirstName())
}
