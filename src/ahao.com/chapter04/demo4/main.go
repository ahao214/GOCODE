package main

import(
	"fmt"
)

//range子语句

func main(){
	str:="abcd"
	for i,char:=range str{
		fmt.Printf("字符串第%d个字符的值为%d\n",i,char)
	}

	for _,char:=range str{
		fmt.Println(char)
	}
	for i:=range str{
		fmt.Println(i)
	}
	for range str{
		fmt.Println("执行成功")
	}

	m:=map[string]int{"a":1,"b":2}
	
	for k,v:=range m{
		fmt.Println(k,v)
	}
	 
	numbers:=[]int{1,2,3,4}
	for i,x:=range numbers{
		fmt.Printf("第%d次，x的值为%d.\n",i,x)
	}


}