package main

import(
	"fmt"
)

//打印金字塔
func print(totalLevel int){
	
	//i表示层数
	for i:=1;i<=totalLevel;i++{
		//打印*前先打印空格
		for k:=1;k<=totalLevel-i;k++{
			fmt.Print(" ")
		}
		//j表示打印多少*
		for j:=1;j<=2*i-1;j++{
			fmt.Print("*")
		}
		fmt.Println() 
	}
}

//九九乘法表
func Jiu(num int){
	for i:=1;i<=num;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%v*%v=%v\t",i,j,i*j)
		}
		fmt.Println()
	}
}


func main(){
	var n int
	fmt.Println("请输入要打印的层数")
	fmt.Scanln(&n)
	print(n)
	Jiu(n)
}