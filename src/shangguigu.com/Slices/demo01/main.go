package main

import(
	"fmt"
)


func main(){
	var intArr[5]int=[...]int{1,22,33,4,5}
	//切片
	slice:=intArr[1:3]

	for i,val:=range slice{
		fmt.Printf("下标:%d 值是：%v\n",i,val)
	}
	fmt.Println("slice的元素个数是：",len(slice))
	fmt.Println("slice的容量是：",cap(slice))
}