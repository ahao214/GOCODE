package main

import(
	"fmt"
)

func main(){

	var slice[] int=make([]int,4,8)
	slice[0]=10
	slice[1]=20
	fmt.Println(slice)
	fmt.Printf("slice len:%d\n",len(slice))
	fmt.Printf("slice cap:%d\n",cap(slice))


	var strSlice[]string=[]string{"Tom","Jack","Kevin"}
	fmt.Println(strSlice)
	fmt.Printf("len:%d\n",len(strSlice))
	fmt.Printf("cap:%d\n",cap(strSlice))
}