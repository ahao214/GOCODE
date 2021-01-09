package main

import(
	"fmt"
)

//max私有函数
func max(a,b int)(maxNum int){
	if a>b{
		return a
	}
	return b
}

//Max公共函数
func Max(a,b int)(maxNum int){
	if a>b{
		maxNum=a
	}
	maxNum=b
	return maxNum
}

func main(){
	x:=1
	y:=2
	z:=3
	var maxNum int
	maxNum=max(x,y)
	fmt.Printf("max(%d,%d)=%d\n",x,y,maxNum)
	maxNum=max(y,z)
	fmt.Printf("max(%d,%d)=%d\n",y,z,maxNum)



}