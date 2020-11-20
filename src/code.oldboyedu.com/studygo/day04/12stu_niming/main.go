package main


import "fmt"

//匿名字段
type person struct{
	string 
	int
}

func main(){
	p1:=person{
		"Tom",
		19,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	

}