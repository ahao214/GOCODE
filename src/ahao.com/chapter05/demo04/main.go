package main

import(
	"fmt"
)

//判断是否是奇数
func isOdd(v int) bool{
	if v%2==0{
		return false
	}
	return true
}

//判断是否是偶数
func isEven(v int) bool{
	if v%2==0{
		return true
	}
	return false
}

//声明一个函数类型
type boolFunc func(int) bool

//函数作为一个参数使用
func filter(slice []int,f boolFunc) []int{
	var result []int
	for _,value:=range slice{
		if f(value){
			result=append(result,value)
		}
	}
	return result
}

//函数作为类型
func main(){
	slice:=[]int{3,1,4,5,9,2}
	fmt.Println("slice=",slice)
	odd:=filter(slice,isOdd)	//函数当作值来传递
	fmt.Println("odd:",odd)
	even:=filter(slice,isEven) 	//函数当作值来传递
	fmt.Println("even:",even)


}

