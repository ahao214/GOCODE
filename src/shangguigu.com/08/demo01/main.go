package main

import(
	"fmt"
)

//对函数返回值命名
func getSumAndSub(one int,two int)(sum int,sub int){
	sum=one+two
	sub=one-two
	return
}

//GO支持可变参数
func getSum(n int,args... int)int{
	sum:=n
	for i:=0;i<len(args);i++{
		sum+=args[i]
	}
	return sum
}

func main(){

	a,b:=getSumAndSub(20,10)
	fmt.Printf("a:%v b:%v\n",a,b)

	res:=getSum(1,23,3,4,-10,24)
	fmt.Println(res)
}