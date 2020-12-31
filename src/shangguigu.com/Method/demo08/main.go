package main

import(
	"fmt"
)

type Student struct{
	name string
	gender string 
	age int
	id int
	score float64
}


func(stu *Student)say() string {
	infoStr:=fmt.Sprintf("Student的信息-name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",stu.name,stu.gender,stu.age,stu.id,stu.score)
	return infoStr
}


func main(){
	var stu=Student{
		name:"Tom",
		gender:"male",
		age:18,
		id:1000,
		score:90.09,
	}
	fmt.Println(stu.say())
	
}