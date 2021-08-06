package main

import (
	"fmt"
)

type Message interface {
	sending()
}

type User struct {
	name  string
	phone string
}

type Admin struct {
	name  string
	phone string
}

func (u *User) sending() {
	fmt.Printf("Sending user phone to %s<%s>\n", u.name, u.phone)
}

func (a *Admin) sending() {
	fmt.Printf("Sending admin phone to %s<%s>\n", a.name, a.phone)
}

//实现多态功能
func main() {
	bill := User{"bill", "bill@email.com"}
	sendMessage(&bill)
	lisa := Admin{"lisa", "lisa@email.com"}
	sendMessage(&lisa)
}

func sendMessage(n Message) {
	n.sending()
}
