package main

import(
	"fmt"
)

//递归

func fbn(n int)int{
	if(n==1 || n==2){
		return 1
	}else{
		return fbn(n-1)+fbn(n-2)
	}
}

func main(){
	res:=fbn(3)
	fmt.Println(res)
	fmt.Println(fbn(4))
	fmt.Println(fbn(5))

}