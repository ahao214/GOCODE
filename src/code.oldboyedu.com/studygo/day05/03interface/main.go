package main

import "fmt"

//接口
type speaker interface{
	speak()
}

//引出接口的实例
type cat struct{

}

type dog struct{

}

type person struct{

}

func(c cat)speak(){
	fmt.Println("喵喵喵")
}

func(d dog)speak(){
	fmt.Println("汪汪汪")
}

func(p person)speak(){
	fmt.Println("啊啊啊")
}

func da(x speaker){
	x.speak()
}


func main(){
	var c cat
	var d dog
	var p person

	da(c)
	da(d)
	da(p)

}