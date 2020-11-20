package main

import "fmt"

//构造函数
type person struct{
	name string
	age int
}

type dog struct{
	name string
}

func newPerson(name string, age int) *person{
	return &person{
		name:name,
		age:age,
	}
}


func newDog(name string) dog{
	return dog{
		name:name,
	}
}

func main(){
	p1:=newPerson("元帅",18)
	p2:=newPerson("小兵",20)
	fmt.Println(p1,p2)

	d1:=newDog("小狗")
	fmt.Println(d1)

}