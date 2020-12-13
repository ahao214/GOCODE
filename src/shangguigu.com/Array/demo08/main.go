package main

import (
	"fmt"
	"math/rand"
	"time"

)


func main(){
	intArr:=[...]int{1,4,6,78,90,7,4,1}
	sum:=0
	avg:=0.0
	for _,value:=range intArr{
		sum+=value
	}
	avg=float64(sum)/float64(len(intArr))
	fmt.Printf("sum is:%v\n",sum)
	fmt.Printf("平均值是：%v\n",avg)

	//随机生成5个数，并将其反转输出
	var intArrB[5] int
	//为了每次生成的随机数不一样，需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<5;i++{
		intArrB[i]=rand.Intn(100) //0<=n<100
	
	}
	fmt.Println("交换前：",intArrB)

	//数字交换
	temp:=0
	for i:=0;i<len(intArrB)/2;i++{
		temp=intArrB[len(intArrB)-1-i]
		intArrB[len(intArrB)-1-i]=intArrB[i]
		intArrB[i]=temp
	}
	fmt.Println("交换后：",intArrB)
	

}