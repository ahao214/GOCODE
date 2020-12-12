package main

import(
	"fmt"
)


func main(){

	var intArr[3] int
	//定义完数组之后，数组的各个元素的默认值都是0
	fmt.Println(intArr)

	//数组的第一个元素的地址，就是数组的地址
	//数组元素的地址是连续的
	fmt.Printf("intArr的地址是：%p\nintArr[0]的地址是：%p\nintArr[1]的地址是：%p",&intArr,&intArr[0],&intArr[1])

}