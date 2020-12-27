package main

import(
	"fmt"
)

type Person struct{
	Name string
}

func(p Person)speak(){
	fmt.Println(p.Name,"是一个好人~")
}

func(p Person)jisuan(){
	res:=0
	for i:=1;i<=1000;i++{
		res+=i
	}
	fmt.Println(p.Name,"计算的结果是：",res)
}

func(p Person)jisuan2(n int){
	res:=0
	for i:=1;i<=n;i++{
		res+=i
	}
	fmt.Println(p.Name,"计算的结果是：",res)
}

func(p Person)getSum(x int,y int)int{
	return x+y
}

func main(){
	var p Person
	p.Name="Jack"
	p.speak()

	p.jisuan()
	
	p.jisuan2(100)

	res:=p.getSum(10,20)
	fmt.Println(res)

}