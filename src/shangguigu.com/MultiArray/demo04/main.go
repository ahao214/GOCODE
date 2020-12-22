package main

import(
	"fmt"
)


func main(){
	var scores [3][5]float64
	for i:=0;i<len(scores);i++{
		for j:=0;j<len(scores[i]);j++{
			fmt.Printf("请输入第%d班第%d个同学的成绩",i+1,j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	totalSum:=0.0
	for i:=0;i<len(scores);i++{
		sum:=0.0
		for j:=0;j<len(scores[i]);j++{
			sum+=scores[i][j]
		}
		totalSum+=sum
		fmt.Printf("第%d班的总分为%v，平均分为%v\t",i+1,sum,sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v，平均分为%v\t",totalSum,totalSum/15)
	
}