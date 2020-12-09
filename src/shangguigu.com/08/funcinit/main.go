package main

import(
	"fmt"
	"shangguigu.com/08/funcinit/utils"
)

var age=test()

//全局变量先被初始化
func test()int{
	fmt.Println("test()...")
	return 90
}

//init函数,通常在init函数中完成初始化工作
func init(){
	fmt.Println("init()...")
}

func main(){
	fmt.Println("main()...")
	fmt.Println("Age=",utils.Age)
	fmt.Println("Name=",utils.Name)

}