package main

import "fmt"

//结构体是值类型
type person struct{
	name string
	gender string
}

func f(x person){
	x.gender="女"	//修改的是副本的gender
}

func f2(x *person){
	//(*x).gender="女"	//根据内存地址找到原变量，修改的就是原变量的值
	x.gender="女"	//语法糖，自动根据指针找对应的变量
}

func main(){
	var p person
	p.name="Tom"
	p.gender="男"
	f(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)


	var p3=person{
		name:"元帅",
		gender:"男",
	}
	fmt.Printf("%#v\n", p3)

	//值的顺序要和声明的结构体定义的顺序一致
	p4:=person{
		"小王子",
		"男",
	}

	fmt.Printf("%#v\n", p4)
}