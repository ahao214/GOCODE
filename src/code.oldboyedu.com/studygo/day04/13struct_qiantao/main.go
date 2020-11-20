package main


import "fmt"

//结构体嵌套
type address struct{
	province string
	city string 
}
type person struct{
	name string
	age int
	addr address
}

type company struct{
	name string 
	address	//匿名嵌套
}




func main(){
	p:=person{
		name:"Tom",
		age:19,
		addr:address{
			province:"广东",
			city:"深圳",
		},
	}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.addr.province)
	fmt.Println(p.addr.city)


	c:=company{
		name:"上海文科",
		address:address{
			province:"上海",
			city:"上海市",
		},
	}
	fmt.Println(c)
	fmt.Println(c.name)
	fmt.Println(c.province)
	fmt.Println(c.city)
	
}