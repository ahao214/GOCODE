package main

import(
	"fmt"
)

type AInterface interface{
	Say()
}

type Stu struct{
	Name string
}

//实现接口
func(stu Stu)Say(){
	fmt.Println("Stu Say()")
}

func main(){
	var stu Stu
	var a AInterface=stu
	a.Say()


}