package main

import(
	"fmt"	
)

func main(){
	str1:="hello北京"
	fmt.Println("str len=",len(str1))

	newStr:= []rune(str1)
	for i:=0;i<len(newStr);i++{
		fmt.Printf("字符=%v\n",newStr[i])
	}


}