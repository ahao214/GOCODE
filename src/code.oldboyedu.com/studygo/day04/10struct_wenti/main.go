package main


import "fmt"

type myInt int

type person struct{
	name string 
	age int 
}

func(i myInt) hello(){
	fmt.Println("这是一个int方法")
}

func main(){
	m:=myInt(11)
	m.hello()

	var p = person{
		name:"Tom",
		age:15,
	}
	fmt.Println(p)

	var p1 = person{
		"Jerry",
		19,
	}
	fmt.Println(p1)

}


func newPerson(name string,age int)person{
	return person{
		name:name,
		age:age,
	}
}