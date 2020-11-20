package main

import "fmt"

//计算n的阶乘
func f(n uint64) uint64{
	if n<=1{
		return 1
	}
	return n*f(n-1)
}

func taijie(n uint64) uint64{
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	return taijie(n-1)+taijie(n-2)
}


//递归
func main(){
	res:=f(5)
	fmt.Println(res)

	tj:=taijie(3)
	fmt.Println(tj)
	

}