package main

import "fmt"

//结构体
type person struct{
	name string
	age int 
	gender string
	hobby []string
}


func main(){
	var p person
	p.name="Tom"
	p.age=23
	p.gender="男"
	p.hobby=[]string{"篮球","跑步","骑行"}
	fmt.Println(p)

	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.gender)


	var s struct{
		x string
		y int
	}
	s.x="hello"
	s.y=100
	fmt.Printf("type:%T value:%v\n",s,s)
}