package model

import(
	"fmt"
)

type person struct{
	Name string
	age int
	sal float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person{
	return &person{
		Name:name,
	}
}

//设置年龄
func (p *person) SetAge(age int){
	if age>0 && age<150{
		p.age=age
	}else{
		fmt.Println("输入的年龄不对...")
	}
}

//获取年龄
func(p *person)GetAge()int{
	return p.age
}

//设置薪水
func(p *person)SetSal(sal float64){
	if sal>=3000&&sal<=30000{
		p.sal=sal
	}else{
		fmt.Println("薪水范围不对...")
	}
}
//获取薪水
func(p *person)GetSal()float64{
	return p.sal
}