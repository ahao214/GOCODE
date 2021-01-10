package main

import(
	"fmt"
)

//匿名函数
func main(){
	func(num int)int{
		sum:=0
		for i:=1;i<=num;i++{
			sum+=i
		}
		fmt.Println(sum)
		return sum
	}(100)	//(100)表示对匿名函数的调用

}