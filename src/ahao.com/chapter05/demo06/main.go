package main

import (
	"fmt"
)

//可变参数
func main(){
	ageArr:=[]int{7,9,3,5,1}
	f1(ageArr...)
}


func f1(arr ...int){
	f2(arr...)
	fmt.Println("")
	f3(arr)
}

func f2(arr ...int){
	for _,char:=range arr{
		fmt.Printf("%d",char)
	}
}

func f3(arr []int){
	for _,char:=range arr{
		fmt.Printf("%d",char)
	}
}