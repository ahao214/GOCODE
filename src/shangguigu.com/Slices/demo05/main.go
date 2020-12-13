package main

import(
	"fmt"
)

func test(slice[]int){
	slice[0]=10
}

func main(){
	var slice=[]int{1,2,3,4,5}
	fmt.Println("slice:",slice)
	test(slice)
	fmt.Println("调用test之后slice:",slice)



}