package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

//匿名字段

func main() {
	//初始化
	s1 := Student{Person{"51mh", "man", 20}, 1, "obj"}
	fmt.Println(s1)

	s2 := Student{Person: Person{"51mh", "man", 20}}
	fmt.Println(s2)

	s3 := Student{Person: Person{name: "51mh"}}
	fmt.Println(s3)
}
