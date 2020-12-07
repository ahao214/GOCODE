package main

import(
	"fmt"
)

//猴子吃桃子
func peach(n int)int{
	if n>10 || n<1{
		fmt.Println("输入的天数不对")
		return 0
	}
	if n==10{
		return 1
	}else{
		return (peach(n+1)+1)*2
	}
}
func main(){
	fmt.Println("第1天的桃子数量：",peach(1))
	fmt.Println("第8天的桃子数量：",peach(8))
	fmt.Println("第9天的桃子数量：",peach(9))
}