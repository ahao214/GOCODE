package main


import(
	"fmt"
)

func main(){

	res1:=func(n int,m int)int{
		return n+m
	}(10,20)

	fmt.Println("res=",res1)


	a:=func(n int,m int)int{
		return n-m
	}

	res2:=a(20,10)
	fmt.Println("res2=",res2)
	
}