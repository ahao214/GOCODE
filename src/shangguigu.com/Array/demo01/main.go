package main

import(
	"fmt"
)


func main(){
	//定义一个数组
	var hens[6] float64
	//给数组赋值
	hens[0]=3.0
	hens[1]=3.0
	hens[2]=6.0
	hens[3]=6.0
	hens[4]=2.0
	hens[5]=4.0

	totalWeight:=0.0
	
	for i:=0;i<len(hens);i++{
		totalWeight+=hens[i]
	}
	avgWeight:=fmt.Sprintf("%.2f",totalWeight/float64(len(hens)))

	fmt.Printf("总重量是:%v\n",totalWeight)
	fmt.Printf("平均重量是:%v",avgWeight)



}