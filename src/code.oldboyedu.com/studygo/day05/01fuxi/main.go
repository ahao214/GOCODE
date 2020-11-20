package main

import "fmt"

//有名字的结构体
type tmp struct{
	x int
	y int
}

type person struct{
		name string
		age int
	}

//构造结构体变量的函数
func newPerson(n string,i int)person{
	return person{
		name:n,
		age:i,
	}
}

//方法
func(p person) dream(str string){
	fmt.Printf("%s的梦想是%s.\n",p.name,str)
}

//指针接收者
//1、需要修改结构体变量的指针
//2、结构体本身比较大
func(p *person)guonian(){
	p.age++
}

func main(){

	var b = tmp{
		1, 
		2,
	}
	fmt.Println(b)

	//匿名结构体
	var a = struct{
		x int
		y int
	}{10,20}
	fmt.Println(a)



	

	var x int
	y:=int8(10)
	fmt.Println(x,y)

	var p1 person	//结构体实例化
	p1.name="Jerry"
	p1.age=20
	p2:=person{"Tom",18} //结构体实例化
	fmt.Println(p1,p2)

	p3:=newPerson("Kevin",32)
	fmt.Println(p3)

	p1.dream("做个咸鱼")
	p2.dream("学好GO语言")
	p3.dream("挣钱")

	fmt.Printf("过年前的年龄是：%d\n",p1.age)
	p1.guonian()
	fmt.Printf("过年后的年龄是：%d",p1.age)
}