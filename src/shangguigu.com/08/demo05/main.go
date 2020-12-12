package main

import(
	"fmt"
)

//defer的使用
func sum(n1 int,n2 int)int{
	//遇见defer的时候，暂时不执行，将defer后面的语句压入到独立的栈(defer栈)
	//当函数执行完毕后，在从defer栈，按照先入后出的方式出栈
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok2 n2=",n2)
	n1++
	n2++
	res:=n1 + n2
	fmt.Println("ok3 res=",res)
	return res
}

func main(){
	res:=sum(10,20)
	fmt.Println("res=",res)
}