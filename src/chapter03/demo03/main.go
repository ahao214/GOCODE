package main

import(
	"fmt"
)

func f1(){
	m:=map[int]func(op int)int{}
	m[1]=func(op int)int{return op}
	m[2]=func(op int)int{return op*op}
	m[3]=func(op int)int{return op*op*op}
	fmt.Println(m[1](2),m[2](2),m[3](2))
}

func f2(){
	mySet:=map[int]bool{}
	mySet[1]=true
	n:=3
	if mySet[n]{
		fmt.Printf("%d is existing",n)
	}else{
		fmt.Printf("%d is not existing",n)
	}


}

func main(){
	//f1()
	f2()
}
