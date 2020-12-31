package main

import(
	"fmt"
	"shangguigu.com/Method/demo11/model"
	
)

func main(){
	stu:=model.Student{
		Name:"Jack",
		Age:19,
	}
	fmt.Println(stu)

	//当结构体是小写字母开头，可以通过工厂模式来解决
	var tea=model.NewTeacher("Alen",28)
	fmt.Println(tea)
	fmt.Println("name=",tea.Name,"age=",tea.GetAge())

}