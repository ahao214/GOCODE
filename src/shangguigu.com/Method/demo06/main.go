package main

import(
	"fmt"
)


type MethodUtils struct{

}

//给MethodUtils编写方法
func(mu MethodUtils)Print(n int, m int){
	for i:=1;i<=n;i++{
		for j:=1;j<=m;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils)Area(width float64,length float64)float64{
	return width*length
}

func main(){
	var mu MethodUtils
	mu.Print(2,3)
	res:=mu.Area(3.4,2.0)
	fmt.Println("面积是：",res)
}