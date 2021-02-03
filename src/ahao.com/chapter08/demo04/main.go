package main

import (
	"fmt"
)

type Employee struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

func main() {
	var emp Employee
	emp.ID = 101
	emp.Name = "Jack"
	emp.Address = "shanghai"
	emp.Phone = "110123123123"
	fmt.Printf("emp ID:%d\n", emp.ID)
	fmt.Printf("emp name:%s\n", emp.Name)
	fmt.Printf("emp address:%s\n", emp.Address)
	fmt.Printf("emp phone:%s\n", emp.Phone)

}
