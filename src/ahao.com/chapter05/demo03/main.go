package main

import(
	"fmt"
)

//函数作为参数
func pipe(ff func()int)int{
	return ff()
}

//FormatFunc定义函数类型
type FormatFunc func(s string,x,y int)string

func format(ff FormatFunc,s string,x,y int)string{
	return ff(s,x,y)
}

func main(){
	s1:=pipe(func() int {return 100})
	s2:=format(func(s string,x,y int)string{
		return fmt.Sprintf(s,x,y)
	},"%d %d",10,20)

	fmt.Println(s1,s2)



}