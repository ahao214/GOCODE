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

func main(){

	a,b:=getSumAndSub(20,10)
	fmt.Printf("a:%v b:%v\n",a,b)

}