package main

import "fmt"

//结构体
type User struct {
	Name  string
	Email string
}

//方法
func (u User) Notify() {
	fmt.Println("%v:%v \n", u.Name, u.Email)
}

func (u *User) Notifyone() {
	fmt.Println("%v:%v \n", u.Name, u.Email)
}

func main() {
	//值类型调用方法
	u1 := User{"golang", "golang@com"}
	u1.Notify()
	//指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()

	u1.Notifyone()
	u3.Notifyone()
}
