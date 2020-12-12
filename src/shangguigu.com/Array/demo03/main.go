package main

import(
	"fmt"
)

func main(){
	var score[5]int
	for i:=0;i<len(score);i++{
		fmt.Printf("请输入第%d个元素的值\n",i+1)
		fmt.Scanln(&score[i])
	}

	for i:=0;i<len(score);i++{
		fmt.Printf("score[%d]=%v\n",i,score[i])	
	}
}