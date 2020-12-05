package main

import(
	"fmt"
)

func main(){
	a:=11
	if a>20{
		fmt.Println("a大于20\n")
	}else if a<10{
		fmt.Println("a小于10\n")
	}else{
		fmt.Println("a大于10\n")
		fmt.Println("a小于20\n")
	}
	fmt.Printf("a的值是：%d\n",a)

}