package main

import(
	"fmt"
)

type A struct{
	Name string
	age int
}

type B struct{
	A
}

func(a *A)Say(){
	fmt.Println("A Say",a.Name,a.age)
}

func(a *A)hello(){
	fmt.Println("A hello",a.Name,a.age)
}



//继承
func main(){
	var b B
	b.A.Name="jack"
	b.A.age=10
	b.A.Say()
	b.A.hello()

	//上面的写法可以简化为下面的写法
	b.Name="kevin"
	b.age=20
	b.Say()
	b.hello()
}