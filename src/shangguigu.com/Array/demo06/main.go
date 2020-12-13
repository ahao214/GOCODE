package main

import(
	"fmt"
)

func testone(arr [3]int){
	arr[0]=88
}

func testtwo(arr *[3]int){
	(*arr)[0]=99
}

func main(){
	arr:=[3]int{11,22,33}
	testone(arr)
	fmt.Println("调用testone之后arr:",arr)
	testtwo(&arr)
	fmt.Println("调用testtwo之后arr:",arr)
	
}