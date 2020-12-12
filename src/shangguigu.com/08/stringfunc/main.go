package main

import(
	"fmt"	
	"strings"
)

func main(){
	str1:="hello北京"
	fmt.Println("str len=",len(str1))

	newStr:= []rune(str1)
	for i:=0;i<len(newStr);i++{
		fmt.Printf("字符=%v\n",newStr[i])
	}

	index:=strings.Index("abcd","a")
	fmt.Println(index)
	lastIndex:=strings.LastIndex("go golang","go")
	fmt.Println(lastIndex)

	strArr:=strings.Split("hello world,go,ok",",")
	fmt.Println(strArr)
	

}