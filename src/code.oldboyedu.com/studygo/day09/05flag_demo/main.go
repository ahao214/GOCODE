package main

import(
	"fmt"
	"flag"
)


//flag 获取命令行参数
func main(){
	//创建一个标志位参数
	// name:=flag.String("name","lph","请输入姓名：")
	// age:=flag.Int("age",90,"请输入真实年龄：")
	// married:=flag.Bool("married",false,"是否结婚：")
	// //使用flag
	// flag.Parse()
	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)


	var name string
	flag.StringVar(&name,"name","李四","请输入姓名")
	flag.Parse()
	fmt.Println(name)



}