package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	//作为值类型
	var personone Person
	personone.firstName = "张"
	personone.lastName = "散散"

	upPerson(&personone)

	fmt.Printf("这个人的名字是 %s %s \n", personone.firstName, personone.lastName)

	//作为指针
	persontwo := new(Person)
	persontwo.firstName = "zhang"
	persontwo.lastName = "sansan"
	upPerson(persontwo)
	fmt.Printf("这个人的名字是 %s %s \n", persontwo.firstName, persontwo.lastName)

	//作为字变量
	personthree := &Person{"jack", "white"}
	upPerson(personthree)
	fmt.Printf("这个人的名字是 %s %s \n", personthree.firstName, personthree.lastName)

}
