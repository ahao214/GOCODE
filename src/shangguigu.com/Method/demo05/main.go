package main

import(
	"fmt"
)

type MethodUtils struct{

}

//给MethodUtils编写方法
func(m MethodUtils) Print(){
	for i:=1;i<=10;i++{
		for j:=1;j<=8;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main(){
	var mu MethodUtils
	mu.Print()

}