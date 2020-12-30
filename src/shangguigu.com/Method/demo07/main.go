package main

import(
	"fmt"
)


type MethodUtils struct{

}


type Calcuator struct{
	Num1 float64
	Num2 float64
}

//判断一个数是偶数还是奇数
func(mu MethodUtils) JudgeNum(num int){
	if num%2==0{
		fmt.Println(num,"数字是偶数")
	}else{
		fmt.Println(num,"数字是奇数")
	}
}

//计算数字
func(cal *Calcuator) GetRes(operator byte)float64{
	res:=0.0
	switch operator{
	case '+':
		res=cal.Num1+cal.Num2
	case '-':
		res=cal.Num1-cal.Num2
	case '*':
		res=cal.Num1*cal.Num2
	case '/':
		res=cal.Num1/cal.Num2
	default:
		fmt.Println("运算符输入有误!")
	}
	return res
} 


func (mu MethodUtils) Print(n int,m int,key string){
	for i:=1;i<=n;i++{
		for j:=1;j<=m;j++{
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func main(){
	var mu MethodUtils
	mu.JudgeNum(10)

	mu.Print(2,4,"&")


	var cal Calcuator
	cal.Num1 = 10.0
	cal.Num2 = 5.0
	res:=cal.GetRes('+')
	fmt.Println("res=",res)
	
}