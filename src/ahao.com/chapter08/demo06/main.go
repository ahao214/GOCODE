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

func PrintEmp(emp *Employee) {
	fmt.Printf("emp id:%d\n", emp.ID)
	fmt.Printf("emp name:%s\n", emp.Name)
	fmt.Printf("emp address:%s\n", emp.Address)
	fmt.Printf("emp phone:%s\n", emp.Phone)
}

func main() {
	var emp Employee
	emp.ID = 101
	emp.Name = "Jack"
	emp.Address = "shanghai"
	emp.Phone = "123123456"
	PrintEmp(&emp)
}
