package main

import "fmt"

func main(){

	//使用make函数创建切片
	s1:= make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d",s1,len(s1),cap(s1))

	//切片的赋值
	s3:=[]int{1,3,5}
	s4:=s3
	fmt.Println(s3,s4)	
	s3[0]=100
	fmt.Println(s3,s4)

	//切片的变量
	for i:=0;i<len(s3);i++{
		fmt.Println(s3[i])
	}
	for i,v:=range s3{
		fmt.Println(i,v)
	}
}