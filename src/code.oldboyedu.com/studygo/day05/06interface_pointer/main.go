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

func main(){
	var a animal
	c1:=cat{"Tom",4}
	c2:=&cat{"假老练",4}

	a=c1
	fmt.Println(a)	
	a=c2
	fmt.Println(a)



}