package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

type Manager struct {
	User
}

func (self *User) ToString() string {
	return fmt.Sprintf("User:%p,%v", self, self)
}

//匿名字段
func main() {
	m := Manager{User{1, "Tom"}}
	fmt.Println("Manager:%p\n", &m)
	fmt.Println(m.ToString())

}
