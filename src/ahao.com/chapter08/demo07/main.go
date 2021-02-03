package main

import "fmt"

type Employee struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

//形式传参
func OperaEmpOne(emp Employee) {
	emp.ID = 100
}

//指针传参
func OperaEmpTwo(emp *Employee) {
	emp.ID = 100
}

func main() {
	var emp Employee
	emp.ID = 101
	emp.Name = "Jack"
	emp.Address = "shanghai"
	emp.Phone = "12310212"
	fmt.Printf("形式传参之前，id:%d\n", emp.ID)
	OperaEmpOne(emp)
	fmt.Printf("形式传参之后：id:%d\n", emp.ID)

	fmt.Printf("指针传参之前：id:%d\n", emp.ID)
	OperaEmpTwo(&emp)
	fmt.Printf("指针传参之后：id:%d\n", emp.ID)

}
