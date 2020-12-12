package main

import(
	"fmt"
)

func main(){
	//初始化数组的方法
	var one[3]int=[3]int{1,2,3}
	fmt.Println("one:",one)
	var two=[3]int{4,5,6}
	fmt.Println("two:",two)
	var three=[...]int{7,8,9}
	fmt.Println("three:",three)

	//使用下标来给数组赋值
	var four=[...]int{2:100,1:200,3:300,0:400}
	fmt.Println("four:",four)

	//使用下标来给数组赋值
	names:=[...]string{1:"Tom",3:"Jerry",0:"Kevin",2:"Cantel"}
	fmt.Println("names:",names)
}