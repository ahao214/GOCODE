package main

import "fmt"

func main(){
	s:= "hello沙河"
	n:= len(s)
	fmt.Println(n)

	for _,c:= range s{
		fmt.Printf("%c\n",c)

	}

	
	

	//字符串修改
	s3 :="红苹果"
	s4:=[]rune(s3)	//把字符串强制转换成rune切片
	s4[0]='黄'
	fmt.Println(string(s4))
}