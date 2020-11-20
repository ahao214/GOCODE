package main

import "fmt"


func main(){
	s1:=[]string{"上海","深圳","北京"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1))

	s1= append(s1,"广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1))

	ss:=[]string{"武汉","杭州","苏州"}
	s1=append(s1,ss...)	//...表示拆分ss
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1))

}