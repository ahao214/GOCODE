package main

import(
	"fmt"
)

type interger int

type Student struct{
	Name string
	Age int
}
	

func(i interger) print(){
	fmt.Println("i=",i)
}

func(i *interger) change(){
	*i=*i+1
}

func(stu *Student) String()string{
	str:=fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}

func main(){
	var i interger=10
	i.print()

	i.change()
	fmt.Println("i=",i)

	stu:=Student{
		Name:"Tom",
		Age:20,
	}
	fmt.Println(&stu)
	

}