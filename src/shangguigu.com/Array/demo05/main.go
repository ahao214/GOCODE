package main

import(
	"fmt"
)

//数组的遍历
func main(){
	//使用下标来给数组赋值
	names:=[...]string{1:"Tom",3:"Jerry",0:"Kevin",2:"Jack"}
	fmt.Println("names:",names)

	for index,value:=range names{
		fmt.Printf("数组的下标是:%d 数组的值是:%v\n",index,value)
	}
}