package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

//结构体
func main() {
	var m manager
	m.age = 12
	m.name = "ahao"
	m.title = "主题"
	fmt.Println(m)
}
