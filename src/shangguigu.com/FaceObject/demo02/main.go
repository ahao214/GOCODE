package main

import(
	"fmt"
)

type Person struct{
	Name string
	Age int
	Scores [3]float64
	ptr *int //指针
	Slice []int //切片
	Map1 map[string] string
}

type Monster struct{
	Name string
	Age int
}

func main(){
	var p1 Person
	
	fmt.Println(p1)

	if p1.ptr==nil{
		fmt.Println("ok1")
	}

	if p1.Slice==nil{
		fmt.Println("ok2")
	}

	if p1.Map1==nil{
		fmt.Println("ok3")
	}

	//使用切片，要先make
	p1.Slice=make([]int,10)
	p1.Slice[0]=100

	p1.Map1=make(map[string]string)
	p1.Map1["key1"]="Tom"

	fmt.Println(p1)

	var monster1 Monster
	monster1.Name="牛魔王"
	monster1.Age=100

	monster2:=monster1
	monster2.Name="青牛精"

	fmt.Println("monster1=",monster1)
	fmt.Println("monster2=",monster2)

}