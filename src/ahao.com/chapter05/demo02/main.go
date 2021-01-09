package main

import(
	"fmt"
)

//函数多返回值
func SumAndProduct(a,b int)(add int,multiplied int){
	add = a+b
	multiplied=a*b
	return
}

//空白标识符"_"的使用
func ThreeValues()(int,int,float32){
	return 5,6,7.8
}

func main(){
	x:=3
	y:=4
	add,multiplied:=SumAndProduct(x,y)
	fmt.Printf("%d+%d=%d\n",x,y,add)
	fmt.Printf("%d*%d=%d\n",x,y,multiplied)

	var i int
	var f float32
	i,_,f=ThreeValues()
	fmt.Printf("The int:%d,the float:%f\n",i,f)


}