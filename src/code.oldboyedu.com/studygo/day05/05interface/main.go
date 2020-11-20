package main

import "fmt"

//接口的实现
type animal interface{
	move()
	eat(string)
}

type cat struct{
	name string
	feet int8
}


func(c cat) move(){
	fmt.Println("走猫步...")
}

func(c cat)eat(food string){
	fmt.Printf("猫吃%s\n",food)
}


type chicken struct{
	feet int8
}

func(c chicken) move(){
	fmt.Println("鸡动")
}

func(c chicken)eat(food string){
	fmt.Printf("吃鸡%s\n",food)
}


func main(){
	var a animal
	bc:=cat{
		name:"淘气",
		feet:4,
	}

	a=bc
	a.eat("小黄鱼")
	fmt.Println(a)


	kfc:=chicken{
		feet:4,
	}

	a = kfc
	fmt.Printf("%T\n",a)
}