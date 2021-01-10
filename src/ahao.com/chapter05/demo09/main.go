package main

import(
	"fmt"
)

func getSeq()func()int{
	i:=0
	return func()int{
		i++
		return i
	}
}

//闭包
func main(){
	//nextNumber为一个函数，i变量自增i并返回
	nextNumber:=getSeq()

	//调用nextNumber函数，i变量自增1并返回
	fmt.Printf("%d",nextNumber())
	fmt.Printf("%d",nextNumber())
	fmt.Printf("%d",nextNumber())

	//创建新的函数nextNumber1,并查看结果
	nextNumber1:=getSeq()
	fmt.Printf("%d",nextNumber1())
	fmt.Printf("%d",nextNumber1())


}